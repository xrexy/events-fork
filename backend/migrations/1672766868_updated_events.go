package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("j2btg39v8wi0zo0")
		if err != nil {
			return err
		}

		collection.ViewRule = types.Pointer("@request.auth.id != null && @request.auth.id = createdBy.id")

		collection.CreateRule = types.Pointer("@request.auth.id != null")

		collection.UpdateRule = types.Pointer("@request.auth.id != null && @request.auth.id = createdBy.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != null && @request.auth.id = createdBy.id")

		// remove
		collection.Schema.RemoveField("kwz2xsvs")

		// remove
		collection.Schema.RemoveField("guxmyog1")

		// remove
		collection.Schema.RemoveField("jhj7y7cv")

		// remove
		collection.Schema.RemoveField("ubdm9k53")

		// update
		edit_title := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "cl3hlih7",
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
		edit_description := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nmixojzl",
			"name": "description",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_description)
		collection.Schema.AddField(edit_description)

		// update
		edit_createdBy := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "dnui3w5x",
			"name": "createdBy",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": true
			}
		}`), edit_createdBy)
		collection.Schema.AddField(edit_createdBy)

		// update
		edit_thumbnail := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "skdsy8vo",
			"name": "thumbnail",
			"type": "file",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"maxSize": 5242880,
				"mimeTypes": [
					"image/jpg",
					"image/jpeg",
					"image/png",
					"image/svg+xml",
					"image/gif"
				],
				"thumbs": []
			}
		}`), edit_thumbnail)
		collection.Schema.AddField(edit_thumbnail)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("j2btg39v8wi0zo0")
		if err != nil {
			return err
		}

		collection.ViewRule = nil

		collection.CreateRule = nil

		collection.UpdateRule = nil

		collection.DeleteRule = nil

		// add
		del_going := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kwz2xsvs",
			"name": "going",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": null,
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false
			}
		}`), del_going)
		collection.Schema.AddField(del_going)

		// add
		del_maybe := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "guxmyog1",
			"name": "maybe",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": null,
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false
			}
		}`), del_maybe)
		collection.Schema.AddField(del_maybe)

		// add
		del_basedOn := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jhj7y7cv",
			"name": "basedOn",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "ntizmtagfpy0v0k",
				"cascadeDelete": false
			}
		}`), del_basedOn)
		collection.Schema.AddField(del_basedOn)

		// add
		del_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ubdm9k53",
			"name": "date",
			"type": "date",
			"required": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_date)
		collection.Schema.AddField(del_date)

		// update
		edit_title := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "cl3hlih7",
			"name": "title",
			"type": "text",
			"required": true,
			"unique": true,
			"options": {
				"min": 8,
				"max": 64,
				"pattern": ""
			}
		}`), edit_title)
		collection.Schema.AddField(edit_title)

		// update
		edit_description := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nmixojzl",
			"name": "description",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": 8,
				"max": 256,
				"pattern": ""
			}
		}`), edit_description)
		collection.Schema.AddField(edit_description)

		// update
		edit_createdBy := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "dnui3w5x",
			"name": "createdBy",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": true
			}
		}`), edit_createdBy)
		collection.Schema.AddField(edit_createdBy)

		// update
		edit_thumbnail := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "skdsy8vo",
			"name": "thumbnail",
			"type": "file",
			"required": true,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"maxSize": 5242880,
				"mimeTypes": [
					"image/jpg",
					"image/jpeg",
					"image/png",
					"image/svg+xml",
					"image/gif"
				],
				"thumbs": []
			}
		}`), edit_thumbnail)
		collection.Schema.AddField(edit_thumbnail)

		return dao.SaveCollection(collection)
	})
}
