package models

import "time"

type SprintRow struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Project    int64     `db:"project"`
	StartedAt  time.Time `db:"started_at"`
	FinishedAt time.Time `db:"finished_at"`
	Status     Status    `db:"status"`
	InBacklog  int64     `db:"in_backlog"`
	InProgress int64     `db:"in_progress"`
	InDone     int64     `db:"in_done"`
	Issues     []int64   `db:"issues"`
}

type Status int32

const (
	StatusUnknown = Status(iota)
	StatusBacklog
	StatusActive
	StatusFinished
)
