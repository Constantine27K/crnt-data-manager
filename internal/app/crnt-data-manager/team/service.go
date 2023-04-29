package team

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	"github.com/Constantine27K/crnt-sdk/pkg/authorization"
)

type Implementation struct {
	team.UnimplementedTeamRegistryServer
	storage    storage.TeamStorage
	validator  validate.Validator
	authorizer authorization.Authorizer
}

func NewService(
	storage storage.TeamStorage,
	validator validate.Validator,
	authorizer authorization.Authorizer,
) *Implementation {
	return &Implementation{
		validator:  validator,
		storage:    storage,
		authorizer: authorizer,
	}
}
