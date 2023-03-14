package crnt_auth_service

import (
	"context"
	"fmt"
	auth "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Service interface {
	Authorize(ctx context.Context, entity string) (*Payload, error)
}

type service struct {
	client auth.AuthClient
}

func NewService(conn *grpc.ClientConn) Service {
	cl := auth.NewAuthClient(conn)
	return &service{
		client: cl,
	}
}

func (s *service) Authorize(ctx context.Context, entity string) (*Payload, error) {
	const (
		header = "Authorization"
	)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no metadata in context")
	}

	headerValues := md.Get(header)
	if len(headerValues) != 1 {
		return nil, fmt.Errorf("no %s header in context", header)
	}

	ctx = metadata.AppendToOutgoingContext(ctx, header, headerValues[0])

	resp, err := s.client.Authorize(ctx, &auth.AuthRequest{Entity: entity})
	if err != nil {
		return nil, err
	}

	return &Payload{
		Username: resp.GetUsername(),
		Role:     resp.GetRole(),
		Team:     resp.GetTeam(),
	}, nil
}
