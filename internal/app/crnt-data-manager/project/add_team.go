package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) AddResponsibleTeam(ctx context.Context, req *desc.ProjectAddTeamRequest) (*desc.ProjectAddTeamResponse, error) {
	projectID, err := i.storage.AddResponsibleTeam(req.GetId(), req.GetTeamId())
	if err != nil {
		log.Error("failed to add team",
			zap.Int64("project", req.GetId()),
			zap.Int64("team", req.GetTeamId()),
			zap.Error(err))
	}

	return &desc.ProjectAddTeamResponse{ProjectId: projectID}, nil
}
