package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/grizlaz/ya-shortener/internal/audit"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/service"
	pb "github.com/grizlaz/ya-shortener/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcShortenerServer struct {
	pb.UnimplementedShortenerServiceServer
	shortener *service.Service
	baseURL   string
	db        *sql.DB
	audit     *audit.Audit
}

func NewGrpcServer(shortener *service.Service, baseURL string, db *sql.DB, audit *audit.Audit) (*grpc.Server, error) {
	shortenerService := &GrpcShortenerServer{
		shortener: shortener,
		baseURL:   baseURL,
		db:        db,
		audit:     audit,
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(WithAuthInterceptor))
	pb.RegisterShortenerServiceServer(s, shortenerService)
	return s, nil
}

func (s *GrpcShortenerServer) ShortenURL(ctx context.Context, request *pb.URLShortenRequest) (*pb.URLShortenResponse, error) {
	url := request.GetUrl()
	if url == "" {
		return nil, status.Error(codes.InvalidArgument, "empty url")
	}
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant get userID")
	}
	shortening, err := s.shortener.Shorten(ctx, url, userID)
	if err != nil {
		if !errors.Is(err, model.ErrConflict) {
			logger.Log.Sugar().Infof("error shortening url %q: %v", url, err)
			return nil, status.Error(codes.Internal, "error shortening url")
		}
	}

	shortURL, err := service.PrependBaseURL(s.baseURL, shortening.ShortURL)
	if err != nil {
		logger.Log.Sugar().Infof("error generating full url for %q: %v", shortening.ShortURL, err)
		return nil, status.Error(codes.Internal, "error generating full url")
	}
	s.audit.Send(model.AuditMessage{
		TS:     time.Now().Unix(),
		Action: "shorten",
		UserID: userID.String(),
		URL:    url,
	})
	response := pb.URLShortenResponse_builder{Result: &shortURL}.Build()
	return response, nil
}

func (s *GrpcShortenerServer) ExpandURL(ctx context.Context, request *pb.URLExpandRequest) (*pb.URLExpandResponse, error) {
	id := request.GetId()
	redirectURL, err := s.shortener.Redirect(ctx, id)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found short url")
	}
	if errors.Is(err, model.ErrURLDeleted) {
		return nil, status.Error(codes.NotFound, "short url deleted")
	}
	if err != nil {
		logger.Log.Sugar().Infof("error getting redirect url for %q: %v", id, err)
		return nil, status.Error(codes.Internal, "error getting redirect url")
	}

	s.audit.Send(model.AuditMessage{
		TS:     time.Now().Unix(),
		Action: "follow",
		URL:    redirectURL,
	})
	response := pb.URLExpandResponse_builder{Result: &redirectURL}.Build()
	return response, nil
}

func (s *GrpcShortenerServer) ListUserURLs(ctx context.Context, empty *emptypb.Empty) (*pb.UserURLsResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant get userID")
	}

	shortenings, err := s.shortener.GetUserUrls(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "cant find user urls")
	}

	responseURLs := make([]*pb.URLData, 0, len(*shortenings))
	for _, v := range *shortenings {
		shortURL, err := service.PrependBaseURL(s.baseURL, v.ShortURL)
		if err != nil {
			logger.Log.Sugar().Infof("error generating full url for %q: %v", v.ShortURL, err)
			return nil, status.Error(codes.Internal, fmt.Sprintf("error generating full url for %q: %v", v.ShortURL, err))
		}
		responseURLs = append(responseURLs, pb.URLData_builder{
			ShortUrl:    &shortURL,
			OriginalUrl: &v.OriginalURL,
		}.Build())
	}

	response := pb.UserURLsResponse_builder{Url: responseURLs}.Build()
	return response, nil
}
