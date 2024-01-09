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
		new_repeatable := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "sie0woo3",
			"name": "repeatable",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_repeatable)
		collection.Schema.AddField(new_repeatable)

		// add
		new_pace := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "khzqxsxj",
			"name": "pace",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_pace)
		collection.Schema.AddField(new_pace)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("20vk5ys2v509j6m")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("sie0woo3")

		// remove
		collection.Schema.RemoveField("khzqxsxj")

		return dao.SaveCollection(collection)
	})
}
