package sprint

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) AddIssue(ctx context.Context, req *desc.AddIssueRequest) (*desc.AddIssueResponse, error) {
	id, err := i.storage.AddIssue(req.GetSprintId(), req.GetIssueId())
	if err != nil {
		log.Error("failed to add an issue into sprint",
			zap.Any("issue", req.GetIssueId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.AddIssueResponse{IssueId: id}, err
}
