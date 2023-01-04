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

		collection, err := dao.FindCollectionByNameOrId("2ftrzh49skq5l07")
		if err != nil {
			return err
		}

		// add
		new_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qjwyspju",
			"name": "type",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": "^(info|success|warning|error)$"
			}
		}`), new_type)
		collection.Schema.AddField(new_type)

		// update
		edit_title := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rgjsvnpj",
			"name": "title",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_title)
		collection.Schema.AddField(edit_title)

		// update
		edit_body := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "pynaicfq",
			"name": "body",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_body)
		collection.Schema.AddField(edit_body)

		// update
		edit_read := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ppqmhkya",
			"name": "read",
			"type": "bool",
			"required": true,
			"unique": false,
			"options": {}
		}`), edit_read)
		collection.Schema.AddField(edit_read)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("2ftrzh49skq5l07")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qjwyspju")

		// update
		edit_title := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rgjsvnpj",
			"name": "title",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_title)
		collection.Schema.AddField(edit_title)

		// update
		edit_body := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "pynaicfq",
			"name": "body",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_body)
		collection.Schema.AddField(edit_body)

		// update
		edit_read := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ppqmhkya",
			"name": "read",
			"type": "bool",
			"required": false,
			"unique": false,
			"options": {}
		}`), edit_read)
		collection.Schema.AddField(edit_read)

		return dao.SaveCollection(collection)
	})
}
