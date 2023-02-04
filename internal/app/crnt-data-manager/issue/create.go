package issue

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateIssue(ctx context.Context, req *desc.IssueCreateRequest) (*desc.IssueCreateResponse, error) {
	err := i.validator.CheckIssue(req.GetIssue())
	if err != nil {
		log.Error("issue is not valid",
			zap.Any("issue", req.GetIssue()),
			zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := i.storage.Add(req.GetIssue())
	if err != nil {
		log.Error("failed to store an issue",
			zap.Any("issue", req.GetIssue()),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueCreateResponse{Id: id}, nil
}
