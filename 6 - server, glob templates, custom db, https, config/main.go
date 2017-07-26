package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	database "./database"

	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
	"github.com/urfave/negroni"
)

type Config struct {
	Development bool
	Domain      string
}

var conf Config

var tmpl *template.Template

func init() {
	//Read config file
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println("Error with config file")
	}
	fmt.Println(conf)
	// Read in templates
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	var resp string
	fmt.Print("Are we on the server: (y/n) ")
	fmt.Scan(&resp)

	if resp == "y" {
		conf.Development = false
	} else {
		conf.Development = true
	}

	//Open Database
	//Builds In Memory Store and AA Trees
	database.OpenAndIndex("database.db")

	//
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

	if conf.Development {
		fmt.Println("Starting server on port 8000")
		http.ListenAndServe(":8000", n)
	} else {
		go normalHTTP(":80")
		log.Fatal(http.Serve(autocert.NewListener("kylepritchard.co.uk", "www.kylepritchard.co.uk"), n))
	}
}

func normalHTTP(port string) {
	http.ListenAndServe(port, http.HandlerFunc(redirectToHttps))
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
	fmt.Println(r)
	http.Redirect(w, r, "https://"+conf.Domain+r.RequestURI, 307)
}

// Homepage
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := database.GetRange("date", true, 0, 0)
	tmpl.ExecuteTemplate(w, "index.gohtml", data)
}

// Post
func post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	slug := p.ByName("slug")
	data := database.GetOne("slug", slug)
	tmpl.ExecuteTemplate(w, "post.gohtml", data)
}

// Add Post Page - '/newpost'
func newPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl.ExecuteTemplate(w, "form.gohtml", nil)
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
