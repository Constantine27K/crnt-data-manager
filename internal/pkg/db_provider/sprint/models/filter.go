package models

import (
	"time"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	sq "github.com/Masterminds/squirrel"
)

type SprintFilter struct {
	IDs            []int64
	Names          []string
	Projects       []int64
	StartedAfter   time.Time
	FinishedBefore time.Time
	Status         Status
}

func NewFilter(sprint *desc.SprintGetRequest) *SprintFilter {
	return &SprintFilter{
		IDs:            sprint.GetIds(),
		Names:          sprint.GetNames(),
		Projects:       sprint.GetProjects(),
		StartedAfter:   sprint.GetStartedAtAfter().AsTime(),
		FinishedBefore: sprint.GetFinishedAtBefore().AsTime(),
	}
}

func (f *SprintFilter) Apply(query sq.SelectBuilder) sq.SelectBuilder {
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

	if len(f.Projects) > 0 {
		query = query.Where(sq.Eq{
			"project": f.Projects,
		})
	}

	if f.StartedAfter.Unix() != 0 {
		query = query.Where(sq.GtOrEq{
			"started_at": f.StartedAfter,
		})
	}

	if f.FinishedBefore.Unix() != 0 {
		query = query.Where(sq.LtOrEq{
			"finished_at": f.FinishedBefore,
		})
	}

	if f.Status != 0 {
		query = query.Where(sq.Eq{
			"status": f.Status,
		})
	}

	return query
}
