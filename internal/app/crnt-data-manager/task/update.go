package task

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Update(ctx context.Context, req *desc.TaskUpdateRequest) (*desc.TaskUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Kostya has not implemented this yet")
}
