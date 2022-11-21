package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	"opsAgent/cmd/opsAgent/api/internal/agent"
	"opsAgent/conf"
	"opsAgent/pkg/api/util"
	pb "opsAgent/pkg/proto/pbgo/opsAgent"
	"opsAgent/pkg/util/log"
)

var listener net.Listener

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if otherHandler == nil {
			grpcServer.ServeHTTP(w, r)
			return
		}
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func StartServer(config conf.Server) error {
	initialize()

	var err error
	listener, err := getListener()
	if err != nil {
		return fmt.Errorf("Unable to create the api server: %v", err)
	}

	err = util.CreateAndSetAuthToken()
	if err != nil {
		return err
	}

	// gRPC server
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(grpcAuth)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(grpcAuth)),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterAgentServer(s, &server{})

	dopts := []grpc.DialOption{grpc.WithInsecure()}

	// starting grpc gateway
	ctx := context.Background()
	gwmux := runtime.NewServeMux()
	err = pb.RegisterAgentHandlerFromEndpoint(
		ctx, gwmux, Addr, dopts)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.Infof("Method: %s, URI: %s, status: %v", values.Method, values.URI, values.Status)
			return nil
		},
	}))

	e.Any("/v1/grpc/*", func(c echo.Context) error {
		gwmux.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// use httpServeMux
	// mux := http.NewServeMux()
	// mux.Handle("/", gwmux)

	srv := &http.Server{
		Addr:    Addr,
		Handler: grpcHandlerFunc(s, e),
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			// Store the connection in the context so requests can reference it if needed
			return context.WithValue(ctx, agent.ConnContextKey, c)
		},
	}

	go srv.Serve(listener)

	log.Infof("start agent on: %v", Addr)
	return nil
}

func StopServer() {
	if listener != nil {
		listener.Close()
	}
}
