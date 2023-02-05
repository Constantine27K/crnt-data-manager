package team

import (
	"context"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetTeams(ctx context.Context, req *desc.TeamGetRequest) (*desc.TeamGetResponse, error) {
	filter := models.NewFilter(req)

	teams, err := i.storage.Get(filter)
	if err != nil {
		log.Error("failed to get teams",
			zap.Any("request", req),
			zap.Error(err))
		return nil, err
	}

	return &desc.TeamGetResponse{Teams: teams}, nil
}
