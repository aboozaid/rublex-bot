package users

import (
	"context"
	"p2ptrader/modules/users/dto"
	"p2ptrader/modules/users/models"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Service interface {
	GetUserIDByTelegramID(ctx context.Context, telegramID string) (*string, error)
	MarkUserAsVerified(ctx context.Context, userID string) error
	CreateUser(ctx context.Context, createParams dto.CreateUserRequestParams) error
}

type service struct {
	repository repo
}

func NewService(app core.App) Service {
	repository := NewRepository(app)
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

func (s service) CreateUser(ctx context.Context, createParams dto.CreateUserRequestParams) error {
	_, err := s.repository.CreateUser(ctx, func(user *models.User) error {
		name := createParams.FirstName + " " + createParams.LastName
		user.SetName(name)
		user.SetFirstName(createParams.FirstName)
		user.SetLastName(createParams.LastName)
		user.SetTelegramID(createParams.TelegramID)
		if createParams.TelegramUsername != nil {
			user.SetTelegramUsername(*createParams.TelegramUsername)
		}
		user.SetLanguageCode(createParams.LanguageCode)

		user.SetPassword(gonanoid.Must(12))

		return nil
	})
	
	return err
}