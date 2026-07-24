package handler

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func WithAuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if _, err := GetUserIDFromContext(ctx); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	token := ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}
	if len(token) == 0 {
		return uuid.Nil, status.Error(codes.Unauthenticated, "missing token")
	}
	userID, err := GetUserIDFromToken(token)
	if err != nil {
		return uuid.Nil, status.Error(codes.InvalidArgument, fmt.Sprintf("error get user ID: %v", err))
	}
	return userID, nil
}
