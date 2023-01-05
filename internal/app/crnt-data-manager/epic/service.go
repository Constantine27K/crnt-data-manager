package epic

import desc "github.com/Constantine27K/crnt-data-manager/pkg/tasks/epic"

type Implementation struct {
	desc.UnimplementedEpicRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
