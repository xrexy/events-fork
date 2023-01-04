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

		// add
		new_maybe := &schema.SchemaField{}
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
		}`), new_maybe)
		collection.Schema.AddField(new_maybe)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("j2btg39v8wi0zo0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("guxmyog1")

		return dao.SaveCollection(collection)
	})
}
