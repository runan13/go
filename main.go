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

func main(){
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleHome)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port,nil))
}