package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {

	var port string
	fmt.Print("Choose a port for server: ")
	fmt.Scan(&port)

	mux := httprouter.New()
	mux.GET("/", index)

	n := negroni.New() // Includes some default middlewares
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	portString := strings.Join([]string{":", port}, "")
	fmt.Println("Starting server listening on port", port)
	http.ListenAndServe(portString, n)
}
