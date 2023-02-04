package issue

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetIssueInfoByID(ctx context.Context, req *desc.IssueInfoGetByIDRequest) (*desc.IssueInfoGetByIDResponse, error) {
	issue, err := i.storage.GetInfoByID(req.GetId())
	if err != nil {
		log.Error("failed to get an issue info by id",
			zap.Int64("id", req.GetId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueInfoGetByIDResponse{IssueInfo: issue}, nil
}
