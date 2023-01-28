package models

import (
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	sq "github.com/Masterminds/squirrel"
)

type IssueFilter struct {
	IDs       []int64
	Name      string
	ParentID  int64
	Type      IssueType
	Author    string
	Assigned  string
	SprintID  int64
	ProjectID int64
	Status    int32
}

func NewFilter(req *desc.IssueGetRequest) *IssueFilter {
	return &IssueFilter{
		IDs:       req.GetIds(),
		Name:      req.GetName(),
		ParentID:  req.GetParentId(),
		Type:      MapFromProtoIssueType(req.GetType()),
		Author:    req.GetAuthor(),
		Assigned:  req.GetAssigned(),
		SprintID:  req.GetSprintId(),
		ProjectID: req.GetProjectId(),
		Status:    MapFromProtoStatus(req.GetStatus()),
	}
}

func (f *IssueFilter) Apply(query sq.SelectBuilder) sq.SelectBuilder {
	if len(f.IDs) > 0 {
		query = query.Where(sq.Eq{
			"id": f.IDs,
		})
	}

	if len(f.Name) > 0 {
		query = query.Where(sq.ILike{
			"name": f.Name,
		})
	}

	if f.Type > 0 {
		query = query.Where(sq.Eq{
			"issue_type": f.Type,
		})
	}

	if len(f.Author) > 0 {
		query = query.Where(sq.ILike{
			"author": f.Author,
		})
	}

	if len(f.Assigned) > 0 {
		query = query.Where(sq.ILike{
			"assigned": f.Assigned,
		})
	}

	if f.SprintID > 0 {
		query = query.Where(sq.Eq{
			"sprint_id": f.SprintID,
		})
	}

	if f.ProjectID > 0 {
		query = query.Where(sq.Eq{
			"project_id": f.ProjectID,
		})
	}

	if f.Status > 0 {
		query = query.Where(sq.Eq{
			"status": f.Status,
		})
	}

	return query
}
