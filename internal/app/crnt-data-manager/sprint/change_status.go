package sprint

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) SprintChangeStatus(ctx context.Context, req *desc.SprintChangeStatusRequest) (*desc.SprintChangeStatusResponse, error) {
	sprintID, err := i.storage.ChangeStatus(req.GetId(), req.GetStatus())
	if err != nil {
		log.Error("failed to change sprint status",
			zap.Any("id", req.GetId()),
			zap.Any("status", req.GetStatus()),
			zap.Error(err))
		return nil, err
	}

	return &desc.SprintChangeStatusResponse{Id: sprintID}, err
}
