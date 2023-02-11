package team

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetTeamByID(ctx context.Context, req *desc.TeamGetByIDRequest) (*desc.TeamGetByIDResponse, error) {
	team, err := i.storage.GetByID(req.GetId())
	if err != nil {
		log.Error("failed to get a team by id",
			zap.Int64("id", req.GetId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.TeamGetByIDResponse{Team: team}, nil
}
