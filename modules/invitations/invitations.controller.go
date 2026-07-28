package invitations

import (
	"net/http"
	"p2ptrader/common"
	"p2ptrader/modules/invitations/dto"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type controller struct {
	service Service
}

func newController(service Service) common.Controller {
	return controller{service}
}

func (c controller) RegisterRoutes(router *router.RouterGroup[*core.RequestEvent]) {
	router.POST("/invitations/verify", c.verifyRequest).
		Unbind(apis.DefaultRequireAuthMiddlewareId)
}

func (c controller) verifyRequest(e *core.RequestEvent) error {
	var requestParams dto.VerifyRequestParams
	if err := e.BindBody(&requestParams); err != nil {
		return err
	}
	if err := requestParams.Validate(); err != nil {
		return err
	}

	err := c.service.VerifyInvitation(e.Request.Context(), requestParams)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, dto.VerifyResponseParams{Verified: true})
}
