package main

import (
	"fmt"
	"html/template"
	"net/http"
	_ "os"
	"time"

	database "./database"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

type Post struct {
	Id         int //'Unique' Key?
	Title      string
	Content    []byte
	Slug       string //slugified version of the title - for routing
	PostDate   time.Time
	FeatureImg string
}

func main() {

	database.OpenAndIndex("database.db")

	// for i := 0; i < 100; i++ {
	// 	database.AddToStore(randomdata.City(), randomdata.Address())
	// }

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
	data := database.GetRange("date", true, 0, 0)
	tpl.Execute(w, data)
}
