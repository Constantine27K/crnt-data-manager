package fixtures

import "github.com/Constantine27K/crnt-data-manager/pkg/api/team"

type teamOpt func(t *team.Team)

func CreateTeam(opts ...teamOpt) *team.Team {
	result := &team.Team{
		Name:          "Team",
		TechOwner:     "T Owner",
		BusinessOwner: "B owner",
		Department:    "Department",
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func WithMembers(names ...string) teamOpt {
	return func(t *team.Team) {
		t.Members = names
	}
}
