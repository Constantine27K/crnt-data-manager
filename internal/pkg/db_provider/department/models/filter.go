package models

import (
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	sq "github.com/Masterminds/squirrel"
)

type DepartmentFilter struct {
	IDs   []int64
	Names []string
}

func NewDepartmentFilter(req *desc.DepartmentGetRequest) *DepartmentFilter {
	return &DepartmentFilter{
		IDs:   req.GetIds(),
		Names: req.GetNames(),
	}
}

func (f *DepartmentFilter) Apply(query sq.SelectBuilder) sq.SelectBuilder {
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

	return query
}
