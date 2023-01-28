package validate

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/template"
	issPack "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
)

type Validator interface {
	CheckIssue(issue *issPack.Issue) error
}

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) CheckIssue(issue *issPack.Issue) error {
	if len(issue.GetName()) == 0 {
		return ErrNameNotFound
	}

	if issue.GetTemplate() == template.Template_TEMPLATE_UNKNOWN {
		return ErrUnknownTemplate
	}

	if issue.GetType() == issPack.IssueType_ISSUE_TYPE_UNKNOWN {
		return ErrUnknownType
	}

	if issue.GetType() == issPack.IssueType_ISSUE_TYPE_SUBTASK &&
		issue.GetParentId() == 0 {
		return ErrNoParentIDSubtask
	}

	if len(issue.GetDescription()) == 0 {
		return ErrNoDescription
	}

	if len(issue.GetAuthor()) == 0 {
		return ErrNoAuthor
	}

	return nil
}
