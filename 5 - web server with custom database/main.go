package main

import (
	"fmt"
	"html/template"
	"net/http"
	_ "os"

	database "github.com/kylepritchard/database"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {

	database.OpenAndIndex("database.db")

	database.AddToStore("Title", []byte("Hello"))

	fmt.Println(database.GetRange("date", false, 0, 0))

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
