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

const userIDName = "userID"

func WithAuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	userID, err := GetUserIDFromMd(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithUserID := context.WithValue(ctx, userIDName, userID)
	return handler(ctxWithUserID, req)
}

func GetUserIDFromMd(ctx context.Context) (uuid.UUID, error) {
	token := ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("authorization")
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

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	v := ctx.Value(userIDName)
	if v == nil {
		return uuid.Nil, fmt.Errorf("not found %s in context", userIDName)
	}
	vUUID, ok := v.(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("cant cast %v to uuid.UUID", v)
	}
	return vUUID, nil
}
