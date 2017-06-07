package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"time"
)

type Post struct {
	Id   string
	Date string
	Name string
}

type Data []Post

func getField(p Post, field string) string {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func (d Data) Sort(field string, asc bool) Data {
	sort.Slice(d, func(i, j int) bool {
		if asc {
			return getField(d[i], field) < getField(d[j], field)
		}
		return getField(d[i], field) > getField(d[j], field)
	})
	return d
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

	var data Data

	ok := json.Unmarshal(b, &data)
	if ok != nil {
		fmt.Printf("Error: %s", ok)
		return
	}

	// data.Ascend()

	// fmt.Println("Name:", data.Sort("Name", true))
	// fmt.Println("Date:", data.Sort("Date", true))
	// fmt.Println("Id:", data.Sort("Id", true))
	fmt.Println("Id:", data.Sort("Id", false))
	fmt.Println("Id:", data.Sort("Id", true))

	elapsed := time.Since(now)
	fmt.Println(elapsed)
}
