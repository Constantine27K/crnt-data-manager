package team

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/rights"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type Implementation struct {
	team.UnimplementedTeamRegistryServer
	storage   storage.TeamStorage
	validator validate.Validator
	verifier  rights.Verifier
}

func NewService(
	storage storage.TeamStorage,
	validator validate.Validator,
	verifier rights.Verifier,
) *Implementation {
	return &Implementation{
		validator: validator,
		storage:   storage,
		verifier:  verifier,
	}
}
