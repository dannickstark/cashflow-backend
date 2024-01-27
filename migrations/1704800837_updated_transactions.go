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

		collection, err := dao.FindCollectionByNameOrId("uo6axjjrwj6yudy")
		if err != nil {
			return err
		}

		// add
		new_next := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "b0dsepcy",
			"name": "next",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "uo6axjjrwj6yudy",
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

		collection, err := dao.FindCollectionByNameOrId("uo6axjjrwj6yudy")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("b0dsepcy")

		return dao.SaveCollection(collection)
	})
}
