package sprint

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) CreateSprint(ctx context.Context, req *desc.SprintCreateRequest) (*desc.SprintCreateResponse, error) {
	id, err := i.storage.Add(req.GetSprint())
	if err != nil {
		log.Error("failed to store an sprint",
			zap.Any("sprint", req.GetSprint()),
			zap.Error(err))
		return nil, err
	}

	return &desc.SprintCreateResponse{Id: id}, err
}
