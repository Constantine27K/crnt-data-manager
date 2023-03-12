package fixtures

import "github.com/Constantine27K/crnt-data-manager/pkg/api/project"

type projOpts func(project *project.Project)

func CreateProject(opts ...projOpts) *project.Project {
	result := &project.Project{
		Name:               "Organization",
		ShortName:          "ORG",
		IsArchived:         false,
		ResponsibleTeamIds: make([]int64, 0),
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func WithName(name string) projOpts {
	return func(project *project.Project) {
		project.Name = name
	}
}

func WithShortName(shortName string) projOpts {
	return func(project *project.Project) {
		project.ShortName = shortName
	}
}

func WithArchived(isArchived bool) projOpts {
	return func(project *project.Project) {
		project.IsArchived = isArchived
	}
}

func WithResponsibleTeams(ids ...int64) projOpts {
	return func(project *project.Project) {
		project.ResponsibleTeamIds = ids
	}
}
