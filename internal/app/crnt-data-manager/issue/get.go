package issue

import (
	"context"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetIssues(ctx context.Context, req *desc.IssueGetRequest) (*desc.IssueGetResponse, error) {
	filter := models.NewFilter(req)

	issues, err := i.storage.Get(filter)
	if err != nil {
		log.Error("failed to get issues",
			zap.Any("request", req),
			zap.Error(err))
		return nil, err
	}

	return &desc.IssueGetResponse{Issues: issues}, nil
}
