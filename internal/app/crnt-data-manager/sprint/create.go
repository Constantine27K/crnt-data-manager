package sprint

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateSprint(ctx context.Context, req *desc.SprintCreateRequest) (*desc.SprintCreateResponse, error) {
	err := i.validator.CheckSprint(req.GetSprint())
	if err != nil {
		log.Error("sprint is not valid",
			zap.Any("sprint", req.GetSprint()),
			zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := i.storage.Add(req.GetSprint())
	if err != nil {
		log.Error("failed to store an sprint",
			zap.Any("sprint", req.GetSprint()),
			zap.Error(err))
		return nil, err
	}

	return &desc.SprintCreateResponse{Id: id}, err
}
