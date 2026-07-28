package users

import (
	"p2ptrader/internal/pocketbase/common"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type controller struct {
	service Service
}

func NewController(db core.App) common.Controller {
	service := NewService(db)
	return controller{service}
}

func (c controller) RegisterRoutes(router *router.RouterGroup[*core.RequestEvent]) {
}
