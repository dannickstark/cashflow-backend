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

		// add
		new_paceUnit := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bg3pnn9f",
			"name": "paceUnit",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"day",
					"week",
					"month",
					"year"
				]
			}
		}`), new_paceUnit)
		collection.Schema.AddField(new_paceUnit)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("20vk5ys2v509j6m")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("bg3pnn9f")

		return dao.SaveCollection(collection)
	})
}
