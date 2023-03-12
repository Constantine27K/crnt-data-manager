package sprint

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) RemoveIssue(ctx context.Context, req *desc.RemoveIssueRequest) (*desc.RemoveIssueResponse, error) {
	sprintID, err := i.storage.RemoveIssue(req.GetId(), req.GetIssueId())
	if err != nil {
		log.Error("failed to remove an issue from sprint",
			zap.Any("issue", req.GetIssueId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.RemoveIssueResponse{SprintId: sprintID}, err
}
