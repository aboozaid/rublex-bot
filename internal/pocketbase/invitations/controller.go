package invitations

import (
	"net/http"
	"p2ptrader/internal/pocketbase/common"
	"p2ptrader/internal/pocketbase/database"
	"p2ptrader/internal/pocketbase/invitations/dto"
	"p2ptrader/internal/pocketbase/users"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type controller struct {
	service Service
}

func NewController(db core.App) common.Controller {
	userService := users.NewService(db)
	dbService := database.NewService(db)
	service := NewService(db, userService, dbService)
	return controller{service}
}

func (c controller) RegisterRoutes(router *router.RouterGroup[*core.RequestEvent]) {
	router.POST("/invitations/verify", c.verifyRequest)
}

func (c controller) verifyRequest(e *core.RequestEvent) error {
	var requestParams dto.VerifyRequestParams
	if err := e.BindBody(&requestParams); err != nil {
		return err
	}
	if err := requestParams.Validate(); err != nil {
		return err
	}

	resp, err := c.service.VerifyInvitation(e.Request.Context(), requestParams)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, resp)
}
