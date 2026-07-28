package common

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type Controller interface {
	RegisterRoutes(router *router.RouterGroup[*core.RequestEvent])
}
