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

		// update
		edit_note := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gvmsugon",
			"name": "note",
			"type": "text",
			"required": false,
			"presentable": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_note)
		collection.Schema.AddField(edit_note)

		// update
		edit_amount := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "06ufmhno",
			"name": "amount",
			"type": "number",
			"required": true,
			"presentable": true,
			"unique": false,
			"options": {
				"min": 0,
				"max": null,
				"noDecimal": false
			}
		}`), edit_amount)
		collection.Schema.AddField(edit_amount)

		// update
		edit_isExpense := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zzptfxpa",
			"name": "isExpense",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_isExpense)
		collection.Schema.AddField(edit_isExpense)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("uo6axjjrwj6yudy")
		if err != nil {
			return err
		}

		// update
		edit_note := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gvmsugon",
			"name": "note",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_note)
		collection.Schema.AddField(edit_note)

		// update
		edit_amount := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "06ufmhno",
			"name": "amount",
			"type": "number",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": 0,
				"max": null,
				"noDecimal": false
			}
		}`), edit_amount)
		collection.Schema.AddField(edit_amount)

		// update
		edit_isExpense := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zzptfxpa",
			"name": "isExpense",
			"type": "bool",
			"required": false,
			"presentable": true,
			"unique": false,
			"options": {}
		}`), edit_isExpense)
		collection.Schema.AddField(edit_isExpense)

		return dao.SaveCollection(collection)
	})
}
