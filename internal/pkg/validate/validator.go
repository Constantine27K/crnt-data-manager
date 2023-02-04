package validate

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/template"
	issPack "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
)

type Validator interface {
	CheckIssue(issue *issPack.Issue) error
}

const (
	maxDescriptionLength = 300
)

type checker func(issue *issPack.Issue) error

type validator struct {
	checks []checker
}

func NewValidator() Validator {
	return &validator{
		checks: []checker{checkTemplate, checkName, checkType, checkDescription, checkAuthor},
	}
}

func (v *validator) CheckIssue(issue *issPack.Issue) error {
	for _, check := range v.checks {
		err := check(issue)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkTemplate(issue *issPack.Issue) error {
	switch issue.GetTemplate() {
	case template.Template_TEMPLATE_COMMON:
		if issue.GetStatus().GetCommon() == nil {
			return ErrNotMatchingTemplate
		}
	case template.Template_TEMPLATE_DEVELOPMENT:
		if issue.GetStatus().GetDevelopment() == nil {
			return ErrNotMatchingTemplate
		}
	case template.Template_TEMPLATE_EPIC:
		if issue.GetStatus().GetEpic() == nil {
			return ErrNotMatchingTemplate
		}
	default:
		return ErrUnknownTemplate
	}

	return nil
}

func checkName(issue *issPack.Issue) error {
	if len(issue.GetName()) == 0 {
		return ErrNameNotFound
	}

	return nil
}

func checkType(issue *issPack.Issue) error {
	if issue.GetType() == issPack.IssueType_ISSUE_TYPE_UNKNOWN {
		return ErrUnknownType
	}

	if issue.GetType() == issPack.IssueType_ISSUE_TYPE_SUBTASK &&
		issue.GetParentId() == 0 {
		return ErrNoParentIDSubtask
	}

	return nil
}

func checkDescription(issue *issPack.Issue) error {
	if len(issue.GetDescription()) == 0 {
		return ErrNoDescription
	}

	if len(issue.GetDescription()) > maxDescriptionLength {
		return ErrTooLongDescription
	}

	return nil
}

func checkAuthor(issue *issPack.Issue) error {
	if len(issue.GetAuthor()) == 0 {
		return ErrNoAuthor
	}

	return nil
}
