package department

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetDepartmentByID(ctx context.Context, req *desc.DepartmentGetByIDRequest) (*desc.DepartmentGetByIDResponse, error) {
	department, err := i.storage.GetByID(req.GetId())
	if err != nil {
		log.Error("failed to get a department by id",
			zap.Int64("id", req.GetId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.DepartmentGetByIDResponse{Department: department}, nil
}
