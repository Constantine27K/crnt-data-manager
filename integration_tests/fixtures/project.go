package fixtures

import "github.com/Constantine27K/crnt-data-manager/pkg/api/project"

type projOpts func(project *project.Project)

func CreateProject(opts ...projOpts) *project.Project {
	result := &project.Project{
		Name:       "Organization",
		ShortName:  "ORG",
		IsArchived: false,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func WithResponsibleTeams(ids ...int64) projOpts {
	return func(project *project.Project) {
		project.ResponsibleTeamIds = ids
	}
}
