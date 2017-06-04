package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl, _ := template.ParseFiles("templates/index.gohtml")
	tpl.Execute(w, nil)

}

func main() {

	var port = "8000"
	// fmt.Print("Choose a port for server: ")
	// fmt.Scan(&port)

	mux := httprouter.New()
	mux.GET("/", index)

	n := negroni.Classic() // Includes some default middlewares
	// n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	portString := strings.Join([]string{":", port}, "")
	fmt.Println("Starting server listening on port", port)
	http.ListenAndServe(portString, n)
}
