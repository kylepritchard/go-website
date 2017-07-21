package main

import (
	"fmt"
	"html/template"
	"net/http"
	_ "os"

	database "./database"

	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
	"github.com/urfave/negroni"
)

// type Post struct {
// 	Id         int //'Unique' Key?
// 	Title      string
// 	Content    template.HTML
// 	Slug       string //slugified version of the title - for routing
// 	PostDate   time.Time
// 	FeatureImg string
// }

func main() {

	database.OpenAndIndex("database.db")

	mux := httprouter.New()

	//Routing
	////////////////////////////////////

	//Homepage
	mux.GET("/", index)

	// Post
	mux.GET("/post/:slug", post)

	// Add new post - GET and POST
	mux.GET("/newpost", newPost)
	mux.POST("/newpost", postToDB)

	//Remove post
	mux.GET("/delete/:slug", removeFromDB)

	n := negroni.Classic() // Log & File Server
	n.UseHandler(mux)

	http.ListenAndServe(":8000", n)

	// log.Fatal(http.Serve(autocert.NewListener("kylepritchard.co.uk"), n))
	fmt.Println("Starting server listening on port 8000")

	//      http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
	//      http.ListenAndServeTLS(":443", "etc/letsencrypt/live/kylepritchard.co.uk/fullchain.pem", "etc/letsencrypt/live/kylepritchard.co.uk/privkey.pem", n)
	//      http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
	http.Redirect(w, r, "https://kylepritchard.co.uk"+r.RequestURI, http.StatusMovedPermanently)
}

// Homepage
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	data := database.GetRange("date", true, 0, 0)
	tpl.Execute(w, data)
}

// Post
func post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tpl := template.Must(template.ParseFiles("templates/post.gohtml"))
	slug := p.ByName("slug")
	data := database.GetOne("slug", slug)
	tpl.Execute(w, data)
}

// Add Post Page - '/newpost'
func newPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl := template.Must(template.ParseFiles("templates/form.gohtml"))
	tpl.Execute(w, nil)
}

// Submit post and save to DB - '/newpost'
func postToDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	markdown := r.PostFormValue("markdown")
	html := blackfriday.MarkdownCommon([]byte(markdown))
	str := string(html)
	database.AddToStore(r.PostFormValue("title"), str)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Remove Post - '/remove/:slug'
func removeFromDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	slug := p.ByName("slug")
	database.Remove(slug)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
