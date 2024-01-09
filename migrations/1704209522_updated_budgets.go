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
		collection.Schema.RemoveField("x1dhzw9w")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("20vk5ys2v509j6m")
		if err != nil {
			return err
		}

		// add
		del_paceUnit := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "x1dhzw9w",
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
		}`), del_paceUnit)
		collection.Schema.AddField(del_paceUnit)

		return dao.SaveCollection(collection)
	})
}
