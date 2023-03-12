package department

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) DepartmentAddProject(ctx context.Context, req *desc.DepartmentAddProjectRequest) (*desc.DepartmentAddProjectResponse, error) {
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
