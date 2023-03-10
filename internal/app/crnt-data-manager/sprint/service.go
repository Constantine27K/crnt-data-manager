package sprint

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
)

type Implementation struct {
	sprint.UnimplementedSprintRegistryServer
	storage   storage.SprintStorage
	validator validate.Validator
}

func NewService(
	storage storage.SprintStorage,
	validator validate.Validator,
) *Implementation {
	return &Implementation{
		storage:   storage,
		validator: validator,
	}
}
