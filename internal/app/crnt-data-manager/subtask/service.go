package subtask

import desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/subtask"

type Implementation struct {
	desc.UnimplementedSubtaskRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
