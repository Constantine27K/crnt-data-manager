package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetProjectByID(ctx context.Context, req *desc.ProjectGetByIDRequest) (*desc.ProjectGetByIDResponse, error) {
	project, err := i.storage.GetByID(req.GetId())
	if err != nil {
		log.Error("failed to get a project",
			zap.Any("id", req.GetId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.ProjectGetByIDResponse{Project: project}, err
}
