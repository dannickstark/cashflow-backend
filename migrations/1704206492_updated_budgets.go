package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("20vk5ys2v509j6m")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("09nkhro6")

		// add
		new_next := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "axfvarwm",
			"name": "next",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "20vk5ys2v509j6m",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_next)
		collection.Schema.AddField(new_next)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("20vk5ys2v509j6m")
		if err != nil {
			return err
		}

		// add
		del_untilDate := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "09nkhro6",
			"name": "untilDate",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_untilDate)
		collection.Schema.AddField(del_untilDate)

		// remove
		collection.Schema.RemoveField("axfvarwm")

		return dao.SaveCollection(collection)
	})
}
