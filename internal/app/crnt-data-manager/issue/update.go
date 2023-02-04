package issue

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) UpdateIssue(ctx context.Context, req *desc.IssueUpdateRequest) (*desc.IssueUpdateResponse, error) {
	id, err := i.storage.Update(req.GetId(), req.GetIssue())
	if err != nil {
		log.Error("failed to update an issue",
			zap.Any("issue", req.GetIssue()),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueUpdateResponse{Id: id}, nil
}
