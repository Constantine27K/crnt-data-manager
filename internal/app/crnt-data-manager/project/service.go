package project

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/storage"
	sprintStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/rights"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type Implementation struct {
	project.UnimplementedProjectRegistryServer
	storage       storage.ProjectStorage
	sprintStorage sprintStorage.SprintStorage
	validator     validate.Validator
	verifier      rights.Verifier
}

func NewService(
	storage storage.ProjectStorage,
	sprintStorage sprintStorage.SprintStorage,
	validator validate.Validator,
	verifier rights.Verifier,
) *Implementation {
	return &Implementation{
		validator:     validator,
		storage:       storage,
		sprintStorage: sprintStorage,
		verifier:      verifier,
	}
}
