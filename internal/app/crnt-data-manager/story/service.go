package story

import desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/story"

type Implementation struct {
	desc.UnimplementedStoryRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
