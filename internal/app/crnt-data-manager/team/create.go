package team

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateTeam(ctx context.Context, req *desc.TeamCreateRequest) (*desc.TeamCreateResponse, error) {
	err := i.verifier.VerifyAdmin(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Any("team", req.GetTeam()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = i.validator.CheckTeam(req.GetTeam())
	if err != nil {
		log.Error("team is not valid",
			zap.Any("team", req.GetTeam()),
			zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := i.storage.Add(req.GetTeam())
	if err != nil {
		log.Error("failed to store a team",
			zap.Any("team", req.GetTeam()),
			zap.Error(err))
		return nil, err
	}

	return &desc.TeamCreateResponse{Id: id}, nil
}
