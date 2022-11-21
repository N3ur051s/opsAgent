package api

import (
	"context"
	"errors"

	"opsAgent/pkg/api/util"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contextKey int

const (
	contextKeyTokenInfoID contextKey = iota
)

var (
	Addr string
)

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

func initialize() {
	var err error
	Addr, err = getAddressPort()
	if err != nil {
		panic("unable to get Agent address and port")
	}
}
