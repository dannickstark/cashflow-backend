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

		// remove
		collection.Schema.RemoveField("qc4lvqie")

		// add
		new_repeatable := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nlzyk2g5",
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
			"id": "ojtvcr6c",
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

		// add
		new_paceUnit := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "mnajtx5s",
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

		collection, err := dao.FindCollectionByNameOrId("uo6axjjrwj6yudy")
		if err != nil {
			return err
		}

		// add
		del_repeat := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qc4lvqie",
			"name": "repeat",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": true
			}
		}`), del_repeat)
		collection.Schema.AddField(del_repeat)

		// remove
		collection.Schema.RemoveField("nlzyk2g5")

		// remove
		collection.Schema.RemoveField("ojtvcr6c")

		// remove
		collection.Schema.RemoveField("mnajtx5s")

		return dao.SaveCollection(collection)
	})
}
