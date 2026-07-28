package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"help": "",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"cascadeDelete": true,
					"collectionId": "_pb_users_auth_",
					"help": "",
					"hidden": false,
					"id": "relation2375276105",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "user",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": true,
					"collectionId": "pbc_1445977019",
					"help": "",
					"hidden": false,
					"id": "relation3543904377",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "exchange",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text3373460893",
					"max": 0,
					"min": 1,
					"name": "api_key",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text3663048549",
					"max": 0,
					"min": 1,
					"name": "api_secret",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text551159581",
					"max": 50,
					"min": 1,
					"name": "tg_group_id",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"help": "",
					"hidden": false,
					"id": "bool458715613",
					"name": "is_active",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "bool"
				},
				{
					"help": "",
					"hidden": false,
					"id": "date1257476049",
					"max": "",
					"min": "",
					"name": "deleted_at",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_2324088501",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_vxocqwxbet` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (\n  ` + "`" + `user` + "`" + `,\n  ` + "`" + `exchange` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_02shuanw6o` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `tg_group_id` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_q9hv0fjlzi` + "`" + ` ON ` + "`" + `accounts` + "`" + ` (` + "`" + `tg_group_id` + "`" + `)"
			],
			"listRule": null,
			"name": "accounts",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2324088501")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
