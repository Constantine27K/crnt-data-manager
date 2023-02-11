package sprint

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetSprintByID(ctx context.Context, req *desc.SprintGetByIDRequest) (*desc.SprintGetByIDResponse, error) {
	sprint, err := i.storage.GetByID(req.GetId())
	if err != nil {
		log.Error("failed to get an sprint by id",
			zap.Int64("id", req.GetId()),
			zap.Error(err))
		return nil, err
	}

	return &desc.SprintGetByIDResponse{Sprint: sprint}, err
}
