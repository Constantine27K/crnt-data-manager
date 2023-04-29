package department

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	"github.com/Constantine27K/crnt-sdk/pkg/authorization"
)

type Implementation struct {
	department.UnimplementedDepartmentRegistryServer
	storage    storage.DepartmentStorage
	validator  validate.Validator
	authorizer authorization.Authorizer
}

func NewService(
	storage storage.DepartmentStorage,
	validator validate.Validator,
	authorizer authorization.Authorizer,
) *Implementation {
	return &Implementation{
		storage:    storage,
		validator:  validator,
		authorizer: authorizer,
	}
}
