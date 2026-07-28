package exchanges

import (
	"net/http"
	"p2ptrader/common"
	"p2ptrader/modules/exchanges/dto"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type controller struct {
	service Service
}

func newController(service Service) common.Controller {
	return controller{service}
}

/*
/exchanges/:id/verifyAccount
*/
func (c controller) RegisterRoutes(router *router.RouterGroup[*core.RequestEvent]) {
	router.POST("/exchanges/{id}/verifyAccount", c.verifyAccountRequest)
}

func (c controller) verifyAccountRequest(e *core.RequestEvent) error {
	var requestParams dto.VerifyAccountRequestParams
	if err := e.BindBody(&requestParams); err != nil {
		return err
	}
	if err := requestParams.Validate(); err != nil {
		return err
	}

	account, err := c.service.VerifyAccount(e.Request.Context(), e.Auth.Id, e.Request.PathValue("id"), requestParams)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, account)
}
