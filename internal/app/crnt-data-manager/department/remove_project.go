package department

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DepartmentRemoveProject(ctx context.Context, req *desc.DepartmentRemoveProjectRequest) (*desc.DepartmentRemoveProjectResponse, error) {
	err := i.verifier.VerifyAdmin(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Any("department", req.GetId()),
			zap.Any("project", req.GetProjectId()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	id, err := i.storage.RemoveProject(req.GetId(), req.GetProjectId())
	if err != nil {
		log.Error("failed to store a department",
			zap.Any("department", req.GetId()),
			zap.Any("project", req.GetProjectId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.DepartmentRemoveProjectResponse{Id: id}, nil
}
