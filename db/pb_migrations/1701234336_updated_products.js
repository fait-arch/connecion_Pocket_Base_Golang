/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("hxh6oonubbj1vvk")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "khrfzcut",
    "name": "product_Price",
    "type": "number",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": 0,
      "max": 20,
      "noDecimal": false
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("hxh6oonubbj1vvk")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "khrfzcut",
    "name": "product_Price",
    "type": "number",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "noDecimal": false
    }
  }))

  return dao.saveCollection(collection)
})
