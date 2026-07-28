package users

import (
	"context"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Service interface {
	GetUserIDByTelegramID(ctx context.Context, telegramID string) (*string, error)
	MarkUserAsVerified(ctx context.Context, userID string) error
}

type service struct {
	repository Repository
}

func NewService(db core.App) Service {
	repository := NewRepository(db)
	return service{repository}
}

func (s service) GetUserIDByTelegramID(ctx context.Context, telegramID string) (*string, error) {
	user, err := s.repository.FindUserByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, apis.NewNotFoundError("USER_NOT_FOUND", err.Error())
	}
	return &user.Id, nil
}

func (s service) MarkUserAsVerified(ctx context.Context, userID string) error {
	if err := s.repository.UpdateUserVerificationStatus(ctx, userID, true); err != nil {
		return apis.NewNotFoundError("USER_NOT_FOUND", err.Error())
	}
	return nil
}
