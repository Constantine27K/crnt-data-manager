package project

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/storage"
	sprintStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	"github.com/Constantine27K/crnt-sdk/pkg/authorization"
)

type Implementation struct {
	project.UnimplementedProjectRegistryServer
	storage       storage.ProjectStorage
	sprintStorage sprintStorage.SprintStorage
	validator     validate.Validator
	authorizer    authorization.Authorizer
}

func NewService(
	storage storage.ProjectStorage,
	sprintStorage sprintStorage.SprintStorage,
	validator validate.Validator,
	authorizer authorization.Authorizer,
) *Implementation {
	return &Implementation{
		validator:     validator,
		storage:       storage,
		sprintStorage: sprintStorage,
		authorizer:    authorizer,
	}
}
