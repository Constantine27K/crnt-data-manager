package issue

import (
	"context"
	"errors"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
)

func (i *Implementation) CreateIssue(ctx context.Context, req *desc.IssueCreateRequest) (*desc.IssueCreateResponse, error) {
	if req.GetIssue() == nil {
		return nil, errors.New("issue should not be nil")
	}

	if len(req.GetIssue().GetName()) == 0 {
		return nil, errors.New("name is too short")
	}

	return &desc.IssueCreateResponse{Id: req.GetIssue().GetId()}, nil
}
