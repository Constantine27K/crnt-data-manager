package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) AddResponsibleTeam(ctx context.Context, req *desc.ProjectAddTeamRequest) (*desc.ProjectAddTeamResponse, error) {
	_, err := i.authorizer.AuthorizeAdmin(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Int64("project", req.GetId()),
			zap.Int64("team", req.GetTeamId()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	projectID, err := i.storage.AddResponsibleTeam(req.GetId(), req.GetTeamId())
	if err != nil {
		log.Error("failed to add team",
			zap.Int64("project", req.GetId()),
			zap.Int64("team", req.GetTeamId()),
			zap.Error(err))
	}

	return &desc.ProjectAddTeamResponse{ProjectId: projectID}, nil
}
