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
					"cascadeDelete": true,
					"collectionId": "pbc_568792081",
					"help": "",
					"hidden": false,
					"id": "relation2069996022",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "payment_method",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
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
			"id": "pbc_592783898",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_ceablb1ozw` + "`" + ` ON ` + "`" + `ads_payment_methods` + "`" + ` (\n  ` + "`" + `ad` + "`" + `,\n  ` + "`" + `payment_method` + "`" + `\n)"
			],
			"listRule": null,
			"name": "ads_payment_methods",
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
		collection, err := app.FindCollectionByNameOrId("pbc_592783898")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
