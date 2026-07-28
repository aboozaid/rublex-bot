package telegramgroups

import (
	"context"
	"p2ptrader/modules/database"
	"p2ptrader/modules/telegram-groups/dto"
	"p2ptrader/modules/telegram-groups/models"

	"github.com/pocketbase/pocketbase/core"
)

type Service interface {
	CreateGroup(ctx context.Context, userID string, createGroupParams dto.CreateGroupRequestParams) error
	IsGroupExist(ctx context.Context, groupID string) (bool, error)
}

type service struct {
	repository repo
	dbService  database.Service
}

func newService(app core.App, dbService database.Service) Service {
	repository := newRepository(app)

	return service{repository, dbService}
}

func (s service) CreateGroup(ctx context.Context, userID string, createGroupParams dto.CreateGroupRequestParams) error {
	return s.dbService.Transaction(ctx, func(ctx context.Context) error {
		groupID, err := s.repository.CreateGroup(ctx, func(g *models.Group) error {
			g.SetGroupID(createGroupParams.GroupID)
			g.SetTitle(createGroupParams.Title)
			g.SetUser(userID)

			return nil
		})
		if err != nil {
			return err
		}
		return s.repository.CreateAccountGroup(ctx, func(ag *models.AccountGroup) error {
			ag.SetGroup(groupID)
			ag.SetAccount(createGroupParams.AccountID)
			return err
		})
	})
}

func (s service) IsGroupExist(ctx context.Context, groupID string) (bool, error) {
	total, err := s.repository.GetTotalGroupsByGroupID(ctx, groupID)
	if err != nil {
		return false, err
	}
	return total > 0, nil
}
