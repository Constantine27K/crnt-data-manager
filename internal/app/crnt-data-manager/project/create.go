package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	crnt_status "github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateProject(ctx context.Context, req *desc.ProjectCreateRequest) (*desc.ProjectCreateResponse, error) {
	err := i.validator.CheckProject(req.GetProject())
	if err != nil {
		log.Error("project is not valid",
			zap.Any("project", req.GetProject()),
			zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := i.storage.Add(req.GetProject())
	if err != nil {
		log.Error("failed to store a project",
			zap.Any("project", req.GetProject()),
			zap.Error(err))
		return nil, err
	}

	sprintID, err := i.sprintStorage.Add(&sprint.Sprint{
		Name:    "Backlog",
		Project: id,
		Status:  crnt_status.Sprint_STATUS_SPRINT_ACTIVE,
	})
	if err != nil {
		log.Error("failed to store a project",
			zap.Any("project", req.GetProject()),
			zap.Error(err))
		return nil, err
	}

	log.Infof("created default backlog sprint with id %v for project: %s", sprintID, req.GetProject().GetName())

	return &desc.ProjectCreateResponse{Id: id}, err
}
