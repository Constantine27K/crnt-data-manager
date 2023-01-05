package task

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.TaskDeleteRequest) (*desc.TaskDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Kostya has not implemented this yet")
}
