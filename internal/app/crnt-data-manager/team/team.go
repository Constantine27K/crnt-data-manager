package team

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type Implementation struct {
	team.UnimplementedTeamRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
