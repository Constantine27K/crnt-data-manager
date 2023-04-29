package team

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	"github.com/Constantine27K/crnt-sdk/pkg/rights"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateTeam(ctx context.Context, req *desc.TeamUpdateRequest) (*desc.TeamUpdateResponse, error) {
	payload, err := i.authorizer.AuthorizeUser(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Any("team", req.GetTeam()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	switch {
	case payload.Role == rights.Admin:
		break
	case payload.Team != req.GetId():
		return nil, status.Error(codes.PermissionDenied, "You have no permission to update information of this team")
	case payload.Role == rights.Employee:
		return nil, status.Error(codes.PermissionDenied, "You have no permission for this")
	}

	id, err := i.storage.Update(req.GetId(), req.GetTeam())
	if err != nil {
		log.Error("failed to update a team",
			zap.Any("team", req.GetTeam()),
			zap.Error(err))
		return nil, err
	}

	return &desc.TeamUpdateResponse{Id: id}, nil
}
