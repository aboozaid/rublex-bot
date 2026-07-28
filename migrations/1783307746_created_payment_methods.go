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
					"id": "text1579384326",
					"max": 255,
					"min": 1,
					"name": "name",
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
					"id": "file3834550803",
					"maxSelect": 0,
					"maxSize": 0,
					"mimeTypes": [
						"image/jpeg",
						"image/png",
						"image/svg+xml",
						"image/gif",
						"image/webp"
					],
					"name": "logo",
					"presentable": false,
					"protected": false,
					"required": false,
					"system": false,
					"thumbs": null,
					"type": "file"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text428357991",
					"max": 255,
					"min": 1,
					"name": "method_id",
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
			"id": "pbc_568792081",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_kd52psgh4k` + "`" + ` ON ` + "`" + `payment_methods` + "`" + ` (\n  ` + "`" + `exchange` + "`" + `,\n  ` + "`" + `method_id` + "`" + `,\n  ` + "`" + `is_active` + "`" + `\n)"
			],
			"listRule": null,
			"name": "payment_methods",
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
		collection, err := app.FindCollectionByNameOrId("pbc_568792081")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
