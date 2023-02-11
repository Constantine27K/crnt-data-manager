package sprint

import (
	"context"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) GetSprint(ctx context.Context, req *desc.SprintGetRequest) (*desc.SprintGetResponse, error) {
	filter := models.NewFilter(req)

	sprints, err := i.storage.Get(filter)
	if err != nil {
		log.Error("failed to get sprint",
			zap.Any("request", req),
			zap.Error(err))
		return nil, err
	}

	return &desc.SprintGetResponse{Sprints: sprints}, err
}
