package main

import (
	"fmt"
	"html/template"
	"net/http"
	_ "os"
	"time"

	database "./database"

	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
	"github.com/urfave/negroni"
)

type Post struct {
	Id         int //'Unique' Key?
	Title      string
	Content    template.HTML
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
	mux.GET("/newpost", newPost)
	mux.POST("/newpost", postToDB)

	n := negroni.Classic() // Log & File Server
	n.UseHandler(mux)

	fmt.Println("Starting server listening on port 8000")
	http.ListenAndServe(":8000", n)
}

// Homepage - '/'
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	data := database.GetRange("date", true, 0, 0)
	fmt.Printf("%T", data[0].Content)
	tpl.Execute(w, data)
}

// Homepage - '/newpost'
func newPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl := template.Must(template.ParseFiles("templates/form.gohtml"))
	tpl.Execute(w, nil)
}

// Homepage - '/newpost'
func postToDB(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// r.ParseForm()
	markdown := r.PostFormValue("markdown")
	// stringSlice := []string{"hello", "bye"}

	html := blackfriday.MarkdownCommon([]byte(markdown))
	str := string(html)
	database.AddToStore(r.PostFormValue("title"), str)
	http.Redirect(w, r, "/", 301)
}
