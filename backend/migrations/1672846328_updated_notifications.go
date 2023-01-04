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
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("2ftrzh49skq5l07")
		if err != nil {
			return err
		}

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
	})
}
