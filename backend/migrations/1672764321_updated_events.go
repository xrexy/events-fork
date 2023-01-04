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

		collection, err := dao.FindCollectionByNameOrId("j2btg39v8wi0zo0")
		if err != nil {
			return err
		}

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
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("j2btg39v8wi0zo0")
		if err != nil {
			return err
		}

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
	})
}
