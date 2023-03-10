package project

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type Implementation struct {
	project.UnimplementedProjectRegistryServer
	storage   storage.ProjectStorage
	validator validate.Validator
}

func NewService(
	storage storage.ProjectStorage,
	validator validate.Validator,
) *Implementation {
	return &Implementation{
		storage:   storage,
		validator: validator,
	}
}
