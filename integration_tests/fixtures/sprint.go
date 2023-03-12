package fixtures

import (
	"time"

	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type optSprint func(sp *sprint.Sprint)

func CreateSprint(opts ...optSprint) *sprint.Sprint {
	result := &sprint.Sprint{
		Name:    "First Sprint",
		Project: 1,
		Status:  status.Sprint_STATUS_SPRINT_BACKLOG,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func WithStatus(stat status.Sprint) optSprint {
	return func(sp *sprint.Sprint) {
		sp.Status = stat
	}
}

func WithDates(startedAt, finishedAt time.Time) optSprint {
	return func(sp *sprint.Sprint) {
		sp.StartedAt = timestamppb.New(startedAt)
		sp.FinishedAt = timestamppb.New(finishedAt)
	}
}
