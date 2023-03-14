package department

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/rights"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/department"
)

type Implementation struct {
	department.UnimplementedDepartmentRegistryServer
	storage   storage.DepartmentStorage
	validator validate.Validator
	verifier  rights.Verifier
}

func NewService(
	storage storage.DepartmentStorage,
	validator validate.Validator,
	verifier rights.Verifier,
) *Implementation {
	return &Implementation{
		storage:   storage,
		validator: validator,
		verifier:  verifier,
	}
}
