package issue

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
)

type Implementation struct {
	desc.UnimplementedIssueRegistryServer
	validator validate.Validator
	storage   storage.IssueStorage
}

func NewService(validator validate.Validator, storage storage.IssueStorage) *Implementation {
	return &Implementation{
		validator: validator,
		storage:   storage,
	}
}
