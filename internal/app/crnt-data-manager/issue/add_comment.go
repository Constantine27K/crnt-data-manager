package issue

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) AddComment(ctx context.Context, req *desc.IssueAddCommentRequest) (*desc.IssueAddCommentResponse, error) {
	id, err := i.storage.AddComment(req.GetId(), req.GetComment())
	if err != nil {
		log.Error("failed to add comment",
			zap.Any("issue", req.GetId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueAddCommentResponse{Id: id}, nil
}
