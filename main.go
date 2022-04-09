package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/runan13/gocoin/blockchain"
)

const (
	port string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template
type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			templates.ExecuteTemplate(w, "add", nil)
		case "POST":
			r.ParseForm()
			data := r.Form.Get("blockData")
			blockchain.GetBlockchain().AddBlock(data)
			http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
		}
}

func main(){
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add", handleAdd)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port,nil))
}