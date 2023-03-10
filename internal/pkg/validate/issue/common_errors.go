package issue

import "errors"

var (
	ErrIssueUnnamed             = errors.New("issue should be named")
	ErrIssueUnknownType         = errors.New("issue should have type")
	ErrIssueUnknownTemplate     = errors.New("issue should have template")
	ErrIssueNoParentIDInSubtask = errors.New("subtask should have parentID")
	ErrIssueNoDescription       = errors.New("issue should have description")
	ErrIssueTooLongDescription  = errors.New("description is too long, maximum length is 300 symbols")
	ErrIssueNoAuthor            = errors.New("issue should have author")
	ErrIssueNoProjectID         = errors.New("issue should have linked to a project")
	ErrIssueNotMatchingTemplate = errors.New("template and status do not match")
)
