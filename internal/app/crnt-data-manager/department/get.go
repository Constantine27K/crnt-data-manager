package department

import (
	"context"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetDepartments(ctx context.Context, req *desc.DepartmentGetRequest) (*desc.DepartmentGetResponse, error) {
	filter := models.NewDepartmentFilter(req)

	departments, err := i.storage.Get(filter)
	if err != nil {
		log.Error("failed to get departments",
			zap.Any("request", req),
			zap.Error(err))
		return nil, err
	}

	return &desc.DepartmentGetResponse{Departments: departments}, nil
}
