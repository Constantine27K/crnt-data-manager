package department

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) DepartmentRemoveProject(ctx context.Context, req *desc.DepartmentRemoveProjectRequest) (*desc.DepartmentRemoveProjectResponse, error) {
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
