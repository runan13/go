package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/runan13/gocoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(w, data)
}

func main(){
	http.HandleFunc("/", handleHome)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port,nil))
}