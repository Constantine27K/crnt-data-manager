package fixtures

import (
	"time"

	"github.com/Constantine27K/crnt-data-manager/pkg/api/comments"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/priority"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/template"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type issueOpt func(issue *desc.Issue)

func CreateIssue(opts ...issueOpt) *desc.Issue {
	const (
		deadline = 24 * time.Hour
		sp       = 3
	)

	result := &desc.Issue{
		Name:        "Test issue",
		Type:        desc.IssueType_ISSUE_TYPE_TASK,
		Description: "test issue",
		Author:      "kibragimov",
		Assigned:    "kibragimov",
		Template:    template.Template_TEMPLATE_DEVELOPMENT,
		CreatedAt:   timestamppb.Now(),
		Deadline:    timestamppb.New(time.Now().Add(deadline)),
		Status: &status.IssueStatus{
			Status: &status.IssueStatus_Development{
				Development: &status.IssueStatusDevelopment{
					Status: status.Development_STATUS_DEVELOPMENT_BACKLOG,
				},
			},
		},
		Priority:    priority.Priority_PRIORITY_MEDIUM,
		StoryPoints: sp,
		Comments:    &comments.Comments{Items: make([]*comments.Comment, 0)},
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func CreateChildIssue(opts ...issueOpt) *desc.Issue {
	const (
		deadline = 24 * time.Hour
		sp       = 1
	)

	result := &desc.Issue{
		Name:        "Child issue",
		Type:        desc.IssueType_ISSUE_TYPE_SUBTASK,
		Description: "child issue",
		Author:      "kibragimov",
		Assigned:    "kibragimov",
		Template:    template.Template_TEMPLATE_DEVELOPMENT,
		CreatedAt:   timestamppb.Now(),
		Deadline:    timestamppb.New(time.Now().Add(deadline)),
		Status: &status.IssueStatus{
			Status: &status.IssueStatus_Development{
				Development: &status.IssueStatusDevelopment{
					Status: status.Development_STATUS_DEVELOPMENT_IN_PROGRESS,
				},
			},
		},
		Priority:    priority.Priority_PRIORITY_LOW,
		StoryPoints: sp,
		Comments:    &comments.Comments{Items: make([]*comments.Comment, 0)},
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func WithComment(author, text string) issueOpt {
	return func(issue *desc.Issue) {
		issue.Comments.Items = append(issue.Comments.Items, &comments.Comment{
			Author:    author,
			Text:      text,
			WrittenAt: timestamppb.Now(),
		})
	}
}

func WithChildren(issue *desc.Issue) issueOpt {
	return func(issue *desc.Issue) {
		issue.Children = append(issue.Children, issue)
	}
}

func WithSprint(id int64) issueOpt {
	return func(issue *desc.Issue) {
		issue.SprintId = id
	}
}

func WithProject(id int64) issueOpt {
	return func(issue *desc.Issue) {
		issue.ProjectId = id
	}
}
