package project

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateProject(ctx context.Context, req *desc.ProjectUpdateRequest) (*desc.ProjectUpdateResponse, error) {
	err := i.verifier.VerifyAdmin(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Any("project", req.GetProject()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	id, err := i.storage.Update(req.GetId(), req.GetProject())
	if err != nil {
		log.Error("failed to update a project",
			zap.Any("team", req.GetProject()),
			zap.Error(err))
		return nil, err
	}

	return &desc.ProjectUpdateResponse{Id: id}, err
}
