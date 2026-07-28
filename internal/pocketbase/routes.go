package pocketbase

import (
	"p2ptrader/internal/pocketbase/common"
	"p2ptrader/internal/pocketbase/invitations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func SetupRoutes(app *pocketbase.PocketBase) {
	controllers := []common.Controller{
		invitations.NewController(app.App),
	}

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		group := se.Router.Group("/api/rublex")
		for _, c := range controllers {
			c.RegisterRoutes(group)
		}
		return se.Next()
	})
}

// next steps invitation code creation via telegram command -> binance module -> account flow creation
