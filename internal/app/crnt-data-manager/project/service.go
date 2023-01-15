package project

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type Implementation struct {
	project.UnimplementedProjectRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
