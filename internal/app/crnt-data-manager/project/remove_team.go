package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) RemoveResponsibleTeam(ctx context.Context, req *desc.ProjectRemoveTeamRequest) (*desc.ProjectRemoveTeamResponse, error) {
	err := i.verifier.VerifyAdmin(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Int64("project", req.GetId()),
			zap.Int64("team", req.GetTeamId()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	projectID, err := i.storage.RemoveResponsibleTeam(req.GetId(), req.GetTeamId())
	if err != nil {
		log.Error("failed to remove team",
			zap.Int64("project", req.GetId()),
			zap.Int64("team", req.GetTeamId()),
			zap.Error(err))
	}

	return &desc.ProjectRemoveTeamResponse{ProjectId: projectID}, nil
}
