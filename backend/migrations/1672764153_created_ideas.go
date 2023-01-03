package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "ntizmtagfpy0v0k",
			"created": "2023-01-03 16:42:33.166Z",
			"updated": "2023-01-03 16:42:33.166Z",
			"name": "ideas",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "8xbtwdv2",
					"name": "title",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "gytuqga0",
					"name": "createdBy",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false
					}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ntizmtagfpy0v0k")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
