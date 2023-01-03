package task

import (
	"context"
	"errors"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/task"
)

func (i *impl) Get(ctx context.Context, req *desc.TaskGetRequest) (*desc.TaskGetResponse, error) {
	return nil, errors.New("unimplemented")
}
