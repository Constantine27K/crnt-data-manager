package department

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DepartmentAddProject(ctx context.Context, req *desc.DepartmentAddProjectRequest) (*desc.DepartmentAddProjectResponse, error) {
	_, err := i.authorizer.AuthorizeAdmin(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Any("department", req.GetId()),
			zap.Any("project", req.GetProjectId()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	id, err := i.storage.AddProject(req.GetId(), req.GetProjectId())
	if err != nil {
		log.Error("failed to store a department",
			zap.Any("department", req.GetId()),
			zap.Any("project", req.GetProjectId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.DepartmentAddProjectResponse{Id: id}, nil
}
