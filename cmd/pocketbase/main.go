package main

import (
	"log/slog"
	"os"
	"p2ptrader/modules"
	"slices"
)

func main() {
	isDev := os.Getenv("ENVIRONMENT") == "development"

	if isDev {
		// to be implemented, restore the database from litestream
	}

	/*app := pocketbase.New()

	pocketbase.SetupHooks(app)

	pocketbase.SetupRoutes(app)

	telegram := telegram.New()

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		if err := telegram.Start(); err != nil {
			slog.Error("Failed to connect to telegram bot", "error", err)
			os.Exit(1)
		}
		return e.Next()
	})

	app.OnTerminate().BindFunc(func(e *core.TerminateEvent) error {
		if err := telegram.Stop(); err != nil {
			return err
		}
		return e.Next()
	})*/
	args := os.Args
	isUpserting := slices.Contains(args, "upsert")
	isMigrating := slices.Contains(args, "migrate")
	if isUpserting {
		if err := modules.Upsert(); err != nil {
			slog.Error("Failed to create superuser account", "error", err)
			os.Exit(1)
		}
	} else if isMigrating {
		if err := modules.Migrate(); err != nil {
			slog.Error("Failed to migrate the database", "error", err)
			os.Exit(1)
		}	
	} else {
		app := modules.NewAppModule()
		if err := app.Start(); err != nil {
			slog.Error("Failed to start application", "error", err)
			os.Exit(1)
		}
	}
}
