package models

import (
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	sq "github.com/Masterminds/squirrel"
)

type TeamFilter struct {
	IDs         []int64
	Names       []string
	Departments []string
}

func NewFilter(req *desc.TeamGetRequest) *TeamFilter {
	return &TeamFilter{
		IDs:         req.GetIds(),
		Names:       req.GetNames(),
		Departments: req.GetDepartments(),
	}
}

func (f *TeamFilter) Apply(query sq.SelectBuilder) sq.SelectBuilder {
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

	if len(f.Departments) > 0 {
		query = query.Where(sq.Eq{
			"departments": f.Departments,
		})
	}

	return query
}
