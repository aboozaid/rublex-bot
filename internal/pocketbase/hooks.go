package pocketbase

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func SetupHooks(app *pocketbase.PocketBase) {
	config := GetConfig()
	if config.Profile == "just" {
		app.OnServe().BindFunc(func(e *core.ServeEvent) error {
			dir, _ := os.Getwd()
			executablePath := fmt.Sprintf("%s/main.exe", dir)
			cmd := exec.Command(executablePath, "superuser", "upsert", config.PocketbaseAdminEmail, config.PocketbaseAdminPassword)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				slog.Error("Superuser upsert failed", "error", err)
			}
			return e.Next()
		})
	}

	app.OnRecordAuthRequest().BindFunc(func(e *core.RecordAuthRequestEvent) error {
		if !e.Record.Verified() {
			return e.UnauthorizedError("UNVERIFIED_USER_NOT_ALLOWED", nil)
		}
		return e.Next()
	})

}
