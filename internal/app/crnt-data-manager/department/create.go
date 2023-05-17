package department

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateDepartment(ctx context.Context, req *desc.DepartmentCreateRequest) (*desc.DepartmentCreateResponse, error) {
	_, err := i.authorizer.AuthorizeAdmin(ctx)
	if err != nil {
		log.Error("error while verifying rights",
			zap.Any("department", req.GetDepartment()),
			zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = i.validator.CheckDepartment(req.GetDepartment())
	if err != nil {
		log.Error("department is not valid",
			zap.Any("department", req.GetDepartment()),
			zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := i.storage.Add(req.GetDepartment())
	if err != nil {
		log.Error("failed to store a department",
			zap.Any("department", req.GetDepartment()),
			zap.Error(err))
		return nil, err
	}

	return &desc.DepartmentCreateResponse{Id: id}, nil
}
