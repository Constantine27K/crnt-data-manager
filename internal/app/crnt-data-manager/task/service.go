package task

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/task"
)

type impl struct {
	task.UnimplementedTaskRegistryServer
}

func NewService() *impl {
	return &impl{}
}
