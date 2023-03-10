package project

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type Implementation struct {
	project.UnimplementedProjectRegistryServer
	validator validate.Validator
	storage   storage.ProjectStorage
}

func NewService(
	validator validate.Validator,
	storage storage.ProjectStorage,
) *Implementation {
	return &Implementation{
		validator: validator,
		storage:   storage,
	}
}
