package common

import (
	"github.com/amarnathcjd/gogram/telegram"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type Module[T any] interface {
	GetService() T
}

type AppLifecycle interface {
	OnAppStart() error
	OnAppStop() error
}

type Controller interface {
	RegisterRoutes(*router.RouterGroup[*core.RequestEvent])
}

type Handler interface {
	RegisterHandlers(*telegram.Client)
}
