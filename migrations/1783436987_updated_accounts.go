package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2324088501")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_vxocqwxbet` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (\n  ` + "`" + `user` + "`" + `,\n  ` + "`" + `exchange` + "`" + `\n)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_q9hv0fjlzi` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `tg_group_id` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2324088501")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_vxocqwxbet` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (\n  ` + "`" + `user` + "`" + `,\n  ` + "`" + `exchange` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_02shuanw6o` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `tg_group_id` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_q9hv0fjlzi` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `tg_group_id` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
