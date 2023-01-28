package issue

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetIssueByID(ctx context.Context, req *desc.IssueGetByIDRequest) (*desc.IssueGetByIDResponse, error) {
	issue, err := i.storage.GetByID(req.GetId())
	if err != nil {
		log.Error("failed to get an issue by id",
			zap.Int64("id", req.GetId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueGetByIDResponse{Issue: issue}, nil
}
