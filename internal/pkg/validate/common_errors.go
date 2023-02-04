package validate

import "errors"

var (
	ErrNameNotFound      = errors.New("issue should be named")
	ErrUnknownType       = errors.New("issue should have type")
	ErrUnknownTemplate   = errors.New("issue should have template")
	ErrNoParentIDSubtask = errors.New("subtask should have parentID")
	ErrNoDescription     = errors.New("issue should have description")
	ErrNoAuthor          = errors.New("issue should have autor")
)
