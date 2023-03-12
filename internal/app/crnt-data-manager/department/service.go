package department

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/department"
)

type Implementation struct {
	department.UnimplementedDepartmentRegistryServer
	validator validate.Validator
	storage   storage.DepartmentStorage
}

func NewService(
	validator validate.Validator,
	storage storage.DepartmentStorage,
) *Implementation {
	return &Implementation{
		validator: validator,
		storage:   storage,
	}
}
