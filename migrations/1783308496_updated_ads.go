package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1911549009")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(8, []byte(`{
			"help": "",
			"hidden": false,
			"id": "bool458715613",
			"name": "is_active",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "bool"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1911549009")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("bool458715613")

		return app.Save(collection)
	})
}
