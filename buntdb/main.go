package main

import (
	"fmt"
	"log"

	"github.com/tidwall/buntdb"
)

func main() {
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.CreateIndex("last_name", "*", buntdb.IndexJSON("name.last"))
	// db.Update(func(tx *buntdb.Tx) error {
	// 	tx.Set("1", `{"name":{"first":"Tom","last":"Johnson"},"age":38}`, nil)
	// 	tx.Set("2", `{"name":{"first":"Janet","last":"Prichard"},"age":47}`, nil)
	// 	tx.Set("3", `{"name":{"first":"Carol","last":"Anderson"},"age":52}`, nil)
	// 	tx.Set("4", `{"name":{"first":"Alan","last":"Cooper"},"age":28}`, nil)
	// 	return nil
	// })

	// db.Update(func(tx *buntdb.Tx) error {
	// 	tx.Delete("user:0:name")
	// 	tx.Delete("user:2:name")
	// 	tx.Delete("user:4:name")
	// 	tx.Delete("user:5:name")
	// 	tx.Delete("user:6:name")
	// 	tx.Delete("user:7:name")
	// 	return nil
	// })

	db.View(func(tx *buntdb.Tx) error {
		fmt.Print(tx.wc)
		tx.Ascend("last_name", func(key, val string) bool {
			fmt.Printf("%s %s\n", key, val)
			return true
		})
		return nil
	})

	db.Shrink()

}
