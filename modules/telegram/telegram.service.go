package telegram

import (
	"context"
	"fmt"
	"p2ptrader/modules/accounts"
	telegramgroups "p2ptrader/modules/telegram-groups"
	telegramgroupsDto "p2ptrader/modules/telegram-groups/dto"
	"strconv"

	"p2ptrader/modules/telegram/dto"
)

type Service interface {
	GetUserAccounts(ctx context.Context, telegramID string) ([]dto.TelegramAccount, error)
	CreateGroup(ctx context.Context, accountID string, createGroupParams dto.CreateGroupParams) error
	IsGroupAlreadyUsed(ctx context.Context, tgGroupID string) (bool, error)
}

type service struct {
	accountsService       accounts.Service
	telegramGroupsService telegramgroups.Service
}

func newService(accountsService accounts.Service, telegramGroupsService telegramgroups.Service) Service {
	return service{accountsService, telegramGroupsService}
}

func (s service) GetUserAccounts(ctx context.Context, telegramID string) ([]dto.TelegramAccount, error) {
	accounts, err := s.accountsService.GetAccountsByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, err
	}
	tgAccounts := make([]dto.TelegramAccount, len(accounts))
	for i, a := range accounts {
		tgAccount := dto.TelegramAccount{ID: a.ID, Account: fmt.Sprintf("%s (%s) - %s", a.Name, a.Nickname, a.Exchange)}
		tgAccounts[i] = tgAccount
	}

	return tgAccounts, nil
}

func (s service) CreateGroup(ctx context.Context, accountID string, createGroupParams dto.CreateGroupParams) error {
	userID, err := s.accountsService.GetUserIDByAccountID(ctx, accountID)
	if err != nil {
		return err
	}
	return s.telegramGroupsService.CreateGroup(ctx, userID, telegramgroupsDto.CreateGroupRequestParams{
		GroupID:   strconv.FormatInt(createGroupParams.GroupID, 10),
		AccountID: accountID,
		Title:     createGroupParams.Title,
	})
}

func (s service) IsGroupAlreadyUsed(ctx context.Context, tgGroupID string) (bool, error) {
	return s.telegramGroupsService.IsGroupExist(ctx, tgGroupID)
}
