package task

import (
	"context"
	"errors"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/task"
)

func (i *impl) Create(ctx context.Context, req *desc.TaskCreateRequest) (*desc.TaskCreateResponse, error) {
	if req.GetTask() == nil {
		return nil, errors.New("task should not be nil")
	}

	if len(req.GetTask().GetName()) == 0 {
		return nil, errors.New("name is too short")
	}

	return &desc.TaskCreateResponse{Id: req.GetTask().GetId()}, nil
}
