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
					"help": "",
					"hidden": false,
					"id": "number416033407",
					"max": null,
					"min": null,
					"name": "min_amount",
					"onlyInt": false,
					"presentable": false,
					"required": true,
					"system": false,
					"type": "number"
				},
				{
					"help": "",
					"hidden": false,
					"id": "number432058571",
					"max": null,
					"min": null,
					"name": "max_amount",
					"onlyInt": false,
					"presentable": false,
					"required": true,
					"system": false,
					"type": "number"
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
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text337201703",
					"max": 0,
					"min": 1,
					"name": "crypto_asset",
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
			"id": "pbc_1911549009",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_bqisrf98v3` + "`" + ` ON ` + "`" + `ads` + "`" + ` (\n  ` + "`" + `account` + "`" + `,\n  ` + "`" + `exchange` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_swntzf1zpd` + "`" + ` ON ` + "`" + `ads` + "`" + ` (\n  ` + "`" + `account` + "`" + `,\n  ` + "`" + `min_amount` + "`" + `,\n  ` + "`" + `max_amount` + "`" + `\n)"
			],
			"listRule": null,
			"name": "ads",
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
		collection, err := app.FindCollectionByNameOrId("pbc_1911549009")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
