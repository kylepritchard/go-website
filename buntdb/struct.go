package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"time"
)

// Post type is a struct of the contents of a post
type Post struct {
	ID   string
	Date string
	Name string
}

// Posts type is a slice of Post structs
type Posts []Post

func getField(p Post, field string) string {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

// Sort method for the
func (p Posts) Sort(field string, asc bool) Posts {
	sort.Slice(p, func(i, j int) bool {
		if asc {
			return getField(p[i], field) < getField(p[j], field)
		}
		return getField(p[i], field) > getField(p[j], field)
	})
	return p
}

func main() {

	now := time.Now()
	day := 24 * time.Hour

	posts := []Post{
		{"1", now.Add(-1 * day).String(), "Kyle"},
		{"2", now.String(), "Alan"},
		{"3", now.Add(1 * day).String(), "Robyn"},
	}

	b, err := json.Marshal(posts)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	
	fmt.Println("Marshalled", b)

	var data Posts

	ok := json.Unmarshal(b, &data)
	if ok != nil {
		fmt.Printf("Error: %s", ok)
		return
	}

	// data.Ascend()

	// fmt.Println("Name:", data.Sort("Name", true))
	// fmt.Println("Date:", data.Sort("Date", true))
	// fmt.Println("Id:", data.Sort("Id", true))
	fmt.Println("ID:", data.Sort("ID", false))
	fmt.Println("ID:", data.Sort("ID", true))

	elapsed := time.Since(now)
	fmt.Println(elapsed)
}
