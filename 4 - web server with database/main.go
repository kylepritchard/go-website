package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"time"
	"encoding/json"
	_ "os"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	
	"github.com/tidwall/buntdb"
	
	_ "./packages"
)

type Comment struct {
	ID		string
	PostID		string
	Comment		string
	PostDate time.Time
}

type Post struct {
	ID		string		`json:"id"`
	Title	string	`json:"title"`
	Content		string		`json:"content"`
	PostDate		time.Time		`json:"postDate"`
}

func main() {
	
	//tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open("database.db")
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	defer db.Close()

	//db.CreateIndex("PostDate", "*", buntdb.IndexJSON("postDate"))
	
	post := Post{
		"43",
		"Title",
		"Content",
		time.Now(),
	}
	
	b, _ := json.Marshal(post)
	//bs := b.(string)
	fmt.Println(string(b))
	
	db.Update(func(tx *buntdb.Tx) error {
			tx.Set(post.ID, string(b) , nil)
	// 	tx.Set("2", `{"name":{"first":"Janet","last":"Prichard"},"age":47}`, nil)
	// 	tx.Set("3", `{"name":{"first":"Carol","last":"Anderson"},"age":52}`, nil)
	// 	tx.Set("4", `{"name":{"first":"Alan","last":"Cooper"},"age":28}`, nil)
	 	return nil
	})

	db.Update(func(tx *buntdb.Tx) error {
		//tx.Delete("9876")
	 	//tx.Delete("43")
	 //	tx.Delete("4563")
	 //	tx.Delete("1234")
	// 	tx.Delete("user:6:name")
	// 	tx.Delete("user:7:name")
	 	return nil
 })



	db.Shrink()
	
	/*
	done := kyledb.InitDB("database.kyledb")
	if done {
		fmt.Println("Initedededeeddededddded")
	} else {
		fmt.Println("Not inited")
	}
	
	db, err := kyledb.Load("database.kyledb")
	if err != nil {
		fmt.Println("shitballs")
	}
	defer db.Close()	
	
	done := db.MakeAndSave()
	if done {
		fmt.Println("Made and entry and saved it")
	}
	*/
	
	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("PostDate", func(key, val string) bool {
			fmt.Printf("%s %s\n", key, val)
			return true
		})
		return nil
	})
	
	//var port = "8000"
	mux := httprouter.New()
	mux.GET("/", index)

	n := negroni.Classic() // Log & File Server
	n.UseHandler(mux)

	fmt.Println("Starting server listening on port 8000")
	http.ListenAndServe(":8000", n)
}



// Homepage - '/'
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tpl.Execute(w, nil)
}
