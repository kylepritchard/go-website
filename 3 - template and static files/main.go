package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	
	//tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	
	var port = "8000"
	mux := httprouter.New()
	mux.GET("/", index)

	n := negroni.Classic() // Log & File Server
	n.UseHandler(mux)

	portString := strings.Join([]string{":", port}, "")
	fmt.Println("Starting server listening on port", port)
	http.ListenAndServe(portString, n)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tpl.Execute(w, nil)

}
