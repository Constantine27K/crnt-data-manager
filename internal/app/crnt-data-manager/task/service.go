package task

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/task"
)

type Implementation struct {
	task.UnimplementedTaskRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
