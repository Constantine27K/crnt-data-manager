package sprint

import "github.com/Constantine27K/crnt-data-manager/pkg/sprint"

type Implementation struct {
	sprint.UnimplementedSprintRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
