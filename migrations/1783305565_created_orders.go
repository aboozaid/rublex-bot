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
					"collectionId": "pbc_2324088501",
					"help": "",
					"hidden": false,
					"id": "relation2100713124",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "account",
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
					"id": "text2376035640",
					"max": 0,
					"min": 1,
					"name": "order_id",
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
					"id": "text3690897523",
					"max": 30,
					"min": 1,
					"name": "fiat_currency",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text360826532",
					"max": 30,
					"min": 1,
					"name": "fiat_symbol",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text337201703",
					"max": 30,
					"min": 1,
					"name": "crypto_asset",
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
					"id": "number1231056614",
					"max": null,
					"min": null,
					"name": "crypto_amount",
					"onlyInt": false,
					"presentable": false,
					"required": true,
					"system": false,
					"type": "number"
				},
				{
					"help": "",
					"hidden": false,
					"id": "number851686407",
					"max": null,
					"min": null,
					"name": "total_fiat_amount",
					"onlyInt": false,
					"presentable": false,
					"required": true,
					"system": false,
					"type": "number"
				},
				{
					"help": "",
					"hidden": false,
					"id": "number3914473387",
					"max": null,
					"min": null,
					"name": "exchange_rate",
					"onlyInt": false,
					"presentable": false,
					"required": true,
					"system": false,
					"type": "number"
				},
				{
					"help": "",
					"hidden": false,
					"id": "bool3961274617",
					"name": "is_buying",
					"presentable": false,
					"required": false,
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
					"help": "",
					"hidden": false,
					"id": "select2063623452",
					"maxSelect": 0,
					"name": "status",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "select",
					"values": [
						"completed",
						"pending",
						"cancelled",
						"disputed"
					]
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
			"id": "pbc_3527180448",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_wb4wa3qslq` + "`" + ` ON ` + "`" + `orders` + "`" + ` (\n  ` + "`" + `account` + "`" + `,\n  ` + "`" + `exchange` + "`" + `,\n  ` + "`" + `is_buying` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_y8qim0wzbm` + "`" + ` ON ` + "`" + `orders` + "`" + ` (\n  ` + "`" + `order_id` + "`" + `,\n  ` + "`" + `exchange` + "`" + `\n)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_85t7xxp7pq` + "`" + ` ON ` + "`" + `orders` + "`" + ` (\n  ` + "`" + `order_id` + "`" + `,\n  ` + "`" + `exchange` + "`" + `\n)"
			],
			"listRule": null,
			"name": "orders",
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
		collection, err := app.FindCollectionByNameOrId("pbc_3527180448")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
