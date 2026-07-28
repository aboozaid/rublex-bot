package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_568792081")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"help": "",
			"hidden": false,
			"id": "number3130834348",
			"max": null,
			"min": 1,
			"name": "total_daily_limit",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"help": "",
			"hidden": false,
			"id": "number530651183",
			"max": null,
			"min": 1,
			"name": "total_monthly_limit",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_568792081")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("number3130834348")

		// remove field
		collection.Fields.RemoveById("number530651183")

		return app.Save(collection)
	})
}
