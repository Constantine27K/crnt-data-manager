package validate

import "errors"

var (
	ErrNameNotFound        = errors.New("issue should be named")
	ErrUnknownType         = errors.New("issue should have type")
	ErrUnknownTemplate     = errors.New("issue should have template")
	ErrNoParentIDSubtask   = errors.New("subtask should have parentID")
	ErrNoDescription       = errors.New("issue should have description")
	ErrTooLongDescription  = errors.New("description is too long, maximum length is 300 symbols")
	ErrNoAuthor            = errors.New("issue should have author")
	ErrNotMatchingTemplate = errors.New("template and status do not match")
)
