package task

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetTaskByID(ctx context.Context, req *desc.TaskGetByIDRequest) (*desc.TaskGetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Kostya has not implemented this yet")
}
