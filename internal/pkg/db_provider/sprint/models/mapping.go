package models

import "github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"

func MapFromProtoStatus(stat status.Sprint) Status {
	switch stat {
	case status.Sprint_STATUS_SPRINT_BACKLOG:
		return StatusBacklog
	case status.Sprint_STATUS_SPRINT_ACTIVE:
		return StatusActive
	case status.Sprint_STATUS_SPRINT_FINISHED:
		return StatusFinished
	default:
		return StatusUnknown
	}
}

func MapToProtoStatus(stat Status) status.Sprint {
	switch stat {
	case StatusBacklog:
		return status.Sprint_STATUS_SPRINT_BACKLOG
	case StatusActive:
		return status.Sprint_STATUS_SPRINT_ACTIVE
	case StatusFinished:
		return status.Sprint_STATUS_SPRINT_FINISHED
	default:
		return status.Sprint_STATUS_SPRINT_UNKNOWN
	}
}
