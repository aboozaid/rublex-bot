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
				"CREATE UNIQUE INDEX ` + "`" + `idx_q9hv0fjlzi` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `tg_group_id` + "`" + `) WHERE 'tg_group_id' != ''",
				"CREATE UNIQUE INDEX ` + "`" + `idx_fy6btqxycr` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `api_key_hashed` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"help": "",
			"hidden": false,
			"id": "text1354711149",
			"max": 0,
			"min": 1,
			"name": "api_key_hashed",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_q9hv0fjlzi` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `tg_group_id` + "`" + `) WHERE 'tg_group_id' != ''"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("text1354711149")

		return app.Save(collection)
	})
}
