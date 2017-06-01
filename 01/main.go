package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {

	mux := httprouter.New()
	mux.GET("/", Index)

	n := negroni.New() // Includes some default middlewares
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	http.ListenAndServe(":8010", n)
}
