package models

import (
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	sq "github.com/Masterminds/squirrel"
)

type ProjectFilter struct {
	IDs              []int64
	Names            []string
	IsArchived       bool
	ResponsibleTeams []int64
}

func NewFilter(req *desc.ProjectGetRequest) *ProjectFilter {
	return &ProjectFilter{
		IDs:              req.GetIds(),
		Names:            req.GetNames(),
		IsArchived:       req.GetIsArchived(),
		ResponsibleTeams: req.GetResponsibleTeamIds(),
	}
}

func (f *ProjectFilter) Apply(query sq.SelectBuilder) sq.SelectBuilder {
	if len(f.IDs) > 0 {
		query = query.Where(sq.Eq{
			"id": f.IDs,
		})
	}

	if len(f.Names) > 0 {
		query = query.Where(sq.Eq{
			"name": f.Names,
		})
	}

	if len(f.ResponsibleTeams) > 0 {
		query = query.Where(sq.Eq{
			"responsible_teams": f.ResponsibleTeams,
		})
	}

	query = query.Where(sq.Eq{
		"is_archived": f.IsArchived,
	})

	return query
}
