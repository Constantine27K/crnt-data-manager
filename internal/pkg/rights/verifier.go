package rights

import (
	"context"
	"fmt"

	authService "github.com/Constantine27K/crnt-data-manager/internal/pkg/services/crnt-auth-service"
)

//go:generate mockery --case=underscore  --recursive --name Verifier
type Verifier interface {
	VerifyAdmin(ctx context.Context) error

	VerifyUpdateTeam(ctx context.Context, team string) error
}

type verifier struct {
	authorizer authService.Service
}

func NewVerifier(authorizer authService.Service) Verifier {
	return &verifier{
		authorizer: authorizer,
	}
}

const (
	admin    = "admin"
	lead     = "lead"
	employee = "employee"
)

func (v *verifier) VerifyAdmin(ctx context.Context) error {
	return v.verifyAdmin(ctx)
}

func (v *verifier) VerifyUpdateTeam(ctx context.Context, team string) error {
	if v.isAdmin(ctx) {
		return nil
	}

	payload, err := v.authorizer.Authorize(ctx, authService.EntityUser)
	if err != nil {
		return err
	}

	if payload.Role == lead && payload.Team == team {
		return nil
	}

	return fmt.Errorf("only team lead of this team can edit")
}

func (v *verifier) verifyAdmin(ctx context.Context) error {
	payload, err := v.authorizer.Authorize(ctx, authService.EntityAdmin)
	if err != nil {
		return err
	}

	if payload.Role != admin {
		return fmt.Errorf("only admin can create a team")
	}

	return nil
}

func (v *verifier) isAdmin(ctx context.Context) bool {
	payload, err := v.authorizer.Authorize(ctx, authService.EntityAdmin)
	if err != nil {
		return false
	}

	return payload.Role == admin
}
