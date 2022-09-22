package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"simpleagent/pkg/api/security"
	"simpleagent/pkg/api/util"
)

type contextKey int

const (
	contextKeyTokenInfoID contextKey = iota
)

var (
	tlsKeyPair  *tls.Certificate
	tlsCertPool *x509.CertPool
	tlsAddr     string
)

func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := util.Validate(w, r); err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func parseToken(token string) (struct{}, error) {
	if token != util.GetAuthToken() {
		return struct{}{}, errors.New("Invalid session token")
	}

	return struct{}{}, nil
}

func grpcAuth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	newCtx := context.WithValue(ctx, contextKeyTokenInfoID, tokenInfo)

	return newCtx, nil
}

func buildSelfSignedKeyPair() ([]byte, []byte) {
	hosts := []string{"127.0.0.1", "localhost", "::1"}
	Addr, err := getAddressPort()
	if err == nil {
		hosts = append(hosts, Addr)
	}
	_, rootCertPEM, rootKey, err := security.GenerateRootCert(hosts, 2048)
	if err != nil {
		return nil, nil
	}

	rootKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
	})

	return rootCertPEM, rootKeyPEM
}

func initializeTLS() {
	cert, key := buildSelfSignedKeyPair()
	if cert == nil {
		panic("unable to generate certificate")
	}
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		panic(err)
	}
	tlsKeyPair = &pair
	tlsCertPool = x509.NewCertPool()
	ok := tlsCertPool.AppendCertsFromPEM(cert)
	if !ok {
		panic("bad certs")
	}

	tlsAddr, err = getAddressPort()
	if err != nil {
		panic("unable to get Agent address and port")
	}
}
