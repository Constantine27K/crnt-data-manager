package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) CreateProject(ctx context.Context, req *desc.ProjectCreateRequest) (*desc.ProjectCreateResponse, error) {
	id, err := i.storage.Add(req.GetProject())
	if err != nil {
		log.Error("failed to store a project",
			zap.Any("project", req.GetProject()),
			zap.Error(err))
		return nil, err
	}

	return &desc.ProjectCreateResponse{Id: id}, err
}
