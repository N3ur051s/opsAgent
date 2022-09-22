package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	"simpleagent/cmd/simpleagent/api/internal/agent"
	"simpleagent/conf"
	"simpleagent/pkg/api/util"
	pb "simpleagent/pkg/proto/pbgo/simpleagent"
	"simpleagent/pkg/util/log"
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
	initializeTLS()

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
		// grpc.Creds(credentials.NewClientTLSFromCert(tlsCertPool, tlsAddr)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(grpcAuth)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(grpcAuth)),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterAgentServer(s, &server{})

	// dcreds := credentials.NewTLS(&tls.Config{
	// 	ServerName: tlsAddr,
	// 	RootCAs:    tlsCertPool,
	// })
	dopts := []grpc.DialOption{grpc.WithInsecure()}

	// starting grpc gateway
	ctx := context.Background()
	gwmux := runtime.NewServeMux()
	err = pb.RegisterAgentHandlerFromEndpoint(
		ctx, gwmux, tlsAddr, dopts)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr: tlsAddr,
		// handle grpc calls directly, falling back to `handler` for non-grpc reqs
		Handler: grpcHandlerFunc(s, gwmux),
		// TLSConfig: &tls.Config{
		// 	Certificates: []tls.Certificate{*tlsKeyPair},
		// 	NextProtos:   []string{"h2"},
		// },
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			// Store the connection in the context so requests can reference it if needed
			return context.WithValue(ctx, agent.ConnContextKey, c)
		},
	}

	// tlsListener := tls.NewListener(listener, srv.TLSConfig)

	go srv.Serve(listener) //nolint:errcheck

	log.Infof("start agent on: %v", tlsAddr)
	return nil
}

// StopServer closes the connection and the server
// stops listening to new commands.
func StopServer() {
	if listener != nil {
		listener.Close()
	}
}

// ServerAddress retruns the server address.
func ServerAddress() *net.TCPAddr {
	return listener.Addr().(*net.TCPAddr)
}
