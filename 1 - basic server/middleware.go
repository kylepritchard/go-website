package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// The type of our middleware consists of the original handler we want to wrap and a message
type Middleware struct {
	next    http.Handler
	message string
}

// Make a constructor for our middleware type since its fields are not exported (in lowercase)
func NewMiddleware(next http.Handler, message string) *Middleware {
	return &Middleware{next: next, message: message}
}

// Our middleware handler
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Start Time
	timeStart := time.RFC3339
	fmt.Print(timeStart)
	//Code
	res := w.(http.ResponseWriter)
	fmt.Print(res.Header.Status())
	// method := r.Method

	// fmt.Printf("Request: %v\n", r)
	m.next.ServeHTTP(w, r)
}

// Our handle function
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	m := NewMiddleware(router, "I'm a middleware")
	log.Fatal(http.ListenAndServe(":8000", m))
}

//[negroni] 2017-06-19T22:02:39Z | 404 |   69.683µs | kylepritchard.co.uk | GET /wp-content/themes/konzept/style.css

// res := rw.(ResponseWriter)
// 	log := LoggerEntry{
// 		StartTime: start.Format(l.dateFormat),
// 		Status:    res.Status(),
// 		Duration:  time.Since(start),
// 		Hostname:  r.Host,
// 		Method:    r.Method,
// 		Path:      r.URL.Path,
// 	}
