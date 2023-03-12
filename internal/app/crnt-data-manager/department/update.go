package department

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) UpdateDepartment(ctx context.Context, req *desc.DepartmentUpdateRequest) (*desc.DepartmentUpdateResponse, error) {
	id, err := i.storage.Update(req.GetId(), req.GetDepartment())
	if err != nil {
		log.Error("failed to update a team",
			zap.Any("department", req.GetDepartment()),
			zap.Error(err))
		return nil, err
	}

	return &desc.DepartmentUpdateResponse{Id: id}, nil
}
