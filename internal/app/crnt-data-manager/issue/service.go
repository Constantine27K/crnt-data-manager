package issue

import (
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
)

type Implementation struct {
	desc.UnimplementedIssueRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
