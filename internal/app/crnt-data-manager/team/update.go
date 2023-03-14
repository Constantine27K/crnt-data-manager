package team

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateTeam(ctx context.Context, req *desc.TeamUpdateRequest) (*desc.TeamUpdateResponse, error) {
	err := i.verifier.VerifyUpdateTeam(ctx, req.GetTeam().GetName())
	if err != nil {
		log.Error("error while verifying rights",
			zap.Any("team", req.GetTeam()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	id, err := i.storage.Update(req.GetId(), req.GetTeam())
	if err != nil {
		log.Error("failed to update a team",
			zap.Any("team", req.GetTeam()),
			zap.Error(err))
		return nil, err
	}

	return &desc.TeamUpdateResponse{Id: id}, nil
}
