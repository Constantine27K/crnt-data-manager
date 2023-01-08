package epic

import desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/epic"

type Implementation struct {
	desc.UnimplementedEpicRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
