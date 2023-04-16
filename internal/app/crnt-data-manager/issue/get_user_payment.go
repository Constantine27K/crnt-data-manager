package issue

import (
	"context"

	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	log "github.com/sirupsen/logrus"
)

func (i *Implementation) GetUserPayment(ctx context.Context, req *desc.IssuePaymentGetRequest) (*desc.IssuePaymentGetResponse, error) {
	userToPayment, err := i.storage.GetUserPayment()
	if err != nil {
		log.Error("failed to get user payment")
		return nil, err
	}

	return &desc.IssuePaymentGetResponse{UserToPayment: userToPayment}, nil
}
