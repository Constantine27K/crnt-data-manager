package sprint

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) UpdateSprint(ctx context.Context, req *desc.SprintUpdateRequest) (*desc.SprintUpdateResponse, error) {
	id, err := i.storage.Update(req.GetId(), req.GetSprint())
	if err != nil {
		log.Error("failed to update a sprint",
			zap.Any("sprint", req.GetSprint()),
			zap.Error(err))
		return nil, err
	}

	return &desc.SprintUpdateResponse{Id: id}, err
}
