package issue

import (
	"context"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetIssueInfo(ctx context.Context, req *desc.IssueInfoGetRequest) (*desc.IssueInfoGetResponse, error) {
	filter := models.NewInfoFilter(req)

	issueInfo, err := i.storage.GetInfo(filter)
	if err != nil {
		log.Error("failed to get issues info",
			zap.Any("request", req),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueInfoGetResponse{IssueInfo: issueInfo}, nil
}
