package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
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

	return &desc.ProjectCreateResponse{Id: id}, err
}
