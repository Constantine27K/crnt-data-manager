package project

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/storage"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type Implementation struct {
	project.UnimplementedProjectRegistryServer
	storage storage.ProjectStorage
}

func NewService(storage storage.ProjectStorage) *Implementation {
	return &Implementation{
		storage: storage,
	}
}
