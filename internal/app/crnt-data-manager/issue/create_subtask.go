package issue

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateSubtask(ctx context.Context, req *desc.IssueCreateSubtaskRequest) (*desc.IssueCreateSubtaskResponse, error) {
	err := i.validator.CheckIssue(req.GetChild())
	if err != nil {
		log.Error("issue is not valid",
			zap.Any("parent", req.GetId()),
			zap.Any("child", req.GetChild()),
			zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := i.storage.CreateSubtask(req.GetId(), req.GetChild())
	if err != nil {
		log.Error("failed to create a subtask",
			zap.Any("parent", req.GetId()),
			zap.Any("child", req.GetChild()),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueCreateSubtaskResponse{Id: id}, nil
}
