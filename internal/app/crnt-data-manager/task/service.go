package task

import (
	desc "github.com/Constantine27K/crnt-data-manager/pkg/tasks/task"
)

type Implementation struct {
	desc.UnimplementedTaskRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
