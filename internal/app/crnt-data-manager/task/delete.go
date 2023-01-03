package task

import (
	"context"
	"errors"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/task"
)

func (i *impl) Delete(ctx context.Context, req *desc.TaskDeleteRequest) (*desc.TaskDeleteResponse, error) {
	return nil, errors.New("unimplemented")
}
