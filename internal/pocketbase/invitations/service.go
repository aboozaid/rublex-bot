package invitations

import (
	"context"
	"p2ptrader/internal/pocketbase/database"
	"p2ptrader/internal/pocketbase/invitations/dto"
	"p2ptrader/internal/pocketbase/users"
	"time"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Service interface {
	VerifyInvitation(ctx context.Context, verifyRequestParams dto.VerifyRequestParams) (dto.VerifyResponseParams, error)
}

type service struct {
	repository  Repository
	userService users.Service
	dbService   *database.Service
}

func NewService(db core.App, userService users.Service, dbService *database.Service) Service {
	repository := NewRepository(db)
	return service{repository, userService, dbService}
}

func (s service) VerifyInvitation(ctx context.Context, verifyRequestParams dto.VerifyRequestParams) (dto.VerifyResponseParams, error) {
	var resp dto.VerifyResponseParams
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
		resp = dto.VerifyResponseParams{Verified: true}
		return nil
	})
	return resp, err
}
