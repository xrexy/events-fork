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
		new_basedOn := &schema.SchemaField{}
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
		}`), new_basedOn)
		collection.Schema.AddField(new_basedOn)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("j2btg39v8wi0zo0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("jhj7y7cv")

		return dao.SaveCollection(collection)
	})
}
