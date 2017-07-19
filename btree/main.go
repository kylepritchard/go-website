package main

import (
	"fmt"
	"go-website/btree/database"
)

func main() {

	database.OpenAndIndex("database.db")

	results := database.GetRange("date", false, 0, 0)
	for _, v := range results {
		fmt.Println(string(v.Content))
	}

}
