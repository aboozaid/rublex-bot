package invitations

import (
	"context"
	"p2ptrader/modules/database"
	"p2ptrader/modules/invitations/dto"
	"p2ptrader/modules/users"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Service interface {
	VerifyInvitation(ctx context.Context, verifyRequestParams dto.VerifyRequestParams) error
	CreateInvitation(ctx context.Context) (string, error)
}

type service struct {
	repository  repo
	userService users.Service
	dbService   database.Service
}

func newService(app core.App, userService users.Service, dbService database.Service) Service {
	repository := NewRepository(app)
	return service{repository, userService, dbService}
}

func (s service) VerifyInvitation(ctx context.Context, verifyRequestParams dto.VerifyRequestParams) error {
	err := s.dbService.Transaction(ctx, func(ctx context.Context) error {
		invitation, err := s.repository.FindInvitationByCode(ctx, verifyRequestParams.Code)
		if err != nil {
			return apis.NewNotFoundError("INVITATION_NOT_FOUND", err.Error())
		}
		due_to := invitation.GetDateTime("due_to")
		if time.Now().After(due_to.Time()) || invitation.GetString("user") != "" {
			return apis.NewBadRequestError("INVALID_OR_USED_INVITATION.", nil)
		}
		userID, err := s.userService.GetUserIDByTelegramID(ctx, verifyRequestParams.TelegramID)
		if err != nil {
			return err
		}
		err = s.repository.UpdateInvitation(ctx, func() (*core.Record, error) {
			invitation.Set("user", userID)
			return invitation, nil
		})
		if err != nil {
			return err
		}
		if err := s.userService.MarkUserAsVerified(ctx, *userID); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s service) CreateInvitation(ctx context.Context) (string, error) {
	var code string
	err := s.repository.CreateInvitation(ctx, func(r *core.Record) error {
		c, err := nanoid.New(12)
		if err != nil {
			return err
		}
		r.Set("code", c)
		r.Set("due_to", types.NowDateTime().Add(24*time.Hour))

		code = c
		return err
	})
	return code, err
}
