package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2324088501")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("text551159581")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"cascadeDelete": true,
			"collectionId": "pbc_116277085",
			"help": "",
			"hidden": false,
			"id": "relation1841317061",
			"maxSelect": 0,
			"minSelect": 0,
			"name": "group",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2324088501")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
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
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation1841317061")

		return app.Save(collection)
	})
}
