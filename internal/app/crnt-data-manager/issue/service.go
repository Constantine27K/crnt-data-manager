package issue

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
)

type Implementation struct {
	desc.UnimplementedIssueRegistryServer
	storage   storage.IssueStorage
	validator validate.Validator
}

func NewService(
	storage storage.IssueStorage,
	validator validate.Validator,
) *Implementation {
	return &Implementation{
		storage:   storage,
		validator: validator,
	}
}
