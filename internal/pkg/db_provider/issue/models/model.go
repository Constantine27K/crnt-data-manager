package models

import (
	"time"
)

type IssueType int64

const (
	IssueTypeUnknown IssueType = iota
	IssueTypeStory
	IssueTypeTask
	IssueTypeSubtask
)

type Template int64

const (
	TemplateUnknown Template = iota
	TemplateCommon
	TemplateDevelopment
)

type Status int32

const (
	StatusUnknown     = 0
	StatusCommon      = 100
	StatusDevelopment = 200
	StatusEpic        = 300
)

type Priority int64

const (
	PriorityUnknown Priority = iota
	PriorityCritical
	PriorityHigh
	PriorityMedium
	PriorityLow
)

type IssueRow struct {
	ID            int64     `db:"id"`
	CompositeName string    `db:"composite_name"`
	Name          string    `db:"name"`
	IssueType     IssueType `db:"issue_type"`
	ParentID      int64     `db:"parent_id"`
	Description   string    `db:"description"`
	Comments      string    `db:"comments"`
	Author        string    `db:"author"`
	Assigned      string    `db:"assigned"`
	QA            string    `db:"qa"`
	Reviewer      string    `db:"reviewer"`
	Template      Template  `db:"template"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	Deadline      time.Time `db:"deadline"`
	Status        int32     `db:"status"`
	Priority      Priority  `db:"priority"`
	SprintID      int64     `db:"sprint_id"`
	ProjectID     int64     `db:"project_id"`
	Components    []int64   `db:"components"`
	StoryPoints   int64     `db:"story_points"`
	Payment       float64   `db:"payment"`
	Children      []int64   `db:"children"`
}

type IssueInfoRow struct {
	ID            int64     `db:"id"`
	CompositeName string    `db:"composite_name"`
	Name          string    `db:"name"`
	IssueType     IssueType `db:"issue_type"`
	Assigned      string    `db:"assigned"`
	Status        int32     `db:"status"`
	Priority      Priority  `db:"priority"`
	StoryPoints   int64     `db:"story_points"`
}
