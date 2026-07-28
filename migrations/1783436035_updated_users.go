package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(13, []byte(`{
			"exceptDomains": null,
			"help": "",
			"hidden": false,
			"id": "url2077391994",
			"name": "photo_url",
			"onlyDomains": null,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "url"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("url2077391994")

		return app.Save(collection)
	})
}
