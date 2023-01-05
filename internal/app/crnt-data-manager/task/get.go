package task

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/tasks/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *desc.TaskGetRequest) (*desc.TaskGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Kostya has not implemented this yet")
}
