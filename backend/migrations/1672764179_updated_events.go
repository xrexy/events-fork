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
		new_going := &schema.SchemaField{}
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
		}`), new_going)
		collection.Schema.AddField(new_going)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("j2btg39v8wi0zo0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("kwz2xsvs")

		return dao.SaveCollection(collection)
	})
}
