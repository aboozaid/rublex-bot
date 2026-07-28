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
					"collectionId": "pbc_3527180448",
					"help": "",
					"hidden": false,
					"id": "relation4113142680",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "order",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": true,
					"collectionId": "pbc_1911549009",
					"help": "",
					"hidden": false,
					"id": "relation2011229528",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "ad",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text119824345",
					"max": 255,
					"min": 1,
					"name": "tg_group_topic_id",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"help": "",
					"hidden": false,
					"id": "json1326724116",
					"maxSize": 0,
					"name": "metadata",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "json"
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
			"id": "pbc_3861817060",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_znv2fvkfp6` + "`" + ` ON ` + "`" + `chats` + "`" + ` (` + "`" + `tg_group_topic_id` + "`" + `)"
			],
			"listRule": null,
			"name": "chats",
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
		collection, err := app.FindCollectionByNameOrId("pbc_3861817060")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
