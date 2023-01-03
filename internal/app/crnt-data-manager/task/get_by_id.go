package task

import (
	"context"
	"errors"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/task"
)

func (i *impl) GetByID(ctx context.Context, req *desc.TaskGetByIDRequest) (*desc.TaskGetByIDResponse, error) {
	return nil, errors.New("unimplemented")
}
