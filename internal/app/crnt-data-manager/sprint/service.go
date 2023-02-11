package sprint

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/storage"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
)

type Implementation struct {
	sprint.UnimplementedSprintRegistryServer
	storage storage.SprintStorage
}

func NewService(storage storage.SprintStorage) *Implementation {
	return &Implementation{
		storage: storage,
	}
}
