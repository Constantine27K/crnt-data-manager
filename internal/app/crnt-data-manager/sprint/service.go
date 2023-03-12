package sprint

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
)

type Implementation struct {
	sprint.UnimplementedSprintRegistryServer
	validator validate.Validator
	storage   storage.SprintStorage
}

func NewService(
	validator validate.Validator,
	storage storage.SprintStorage,
) *Implementation {
	return &Implementation{
		validator: validator,
		storage:   storage,
	}
}
