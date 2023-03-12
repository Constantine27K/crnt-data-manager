package team

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type Implementation struct {
	team.UnimplementedTeamRegistryServer
	validator validate.Validator
	storage   storage.TeamStorage
}

func NewService(
	validator validate.Validator,
	storage storage.TeamStorage,
) *Implementation {
	return &Implementation{
		validator: validator,
		storage:   storage,
	}
}
