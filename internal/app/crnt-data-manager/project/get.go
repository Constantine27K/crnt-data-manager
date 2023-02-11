package project

import (
	"context"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetProjects(ctx context.Context, req *desc.ProjectGetRequest) (*desc.ProjectGetResponse, error) {
	filter := models.NewFilter(req)

	projects, err := i.storage.Get(filter)
	if err != nil {
		log.Error("failed to get projects",
			zap.Any("filter", filter),
			zap.Error(err))
		return nil, err
	}

	return &desc.ProjectGetResponse{Projects: projects}, err
}
