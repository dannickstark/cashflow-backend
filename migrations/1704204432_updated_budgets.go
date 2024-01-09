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
		collection.Schema.RemoveField("qxra2vzk")

		// add
		new_untilDate := &schema.SchemaField{}
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
		}`), new_untilDate)
		collection.Schema.AddField(new_untilDate)

		// update
		edit_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "yydmxxra",
			"name": "date",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_date)
		collection.Schema.AddField(edit_date)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("20vk5ys2v509j6m")
		if err != nil {
			return err
		}

		// add
		del_previousStates := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qxra2vzk",
			"name": "previousStates",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_previousStates)
		collection.Schema.AddField(del_previousStates)

		// remove
		collection.Schema.RemoveField("09nkhro6")

		// update
		edit_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "yydmxxra",
			"name": "lastModifiedMonth",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_date)
		collection.Schema.AddField(edit_date)

		return dao.SaveCollection(collection)
	})
}
