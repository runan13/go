package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/runan13/gocoin/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func handleDoc(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Printf("%s", b)
}

func main() {
	http.HandleFunc("/", handleDoc)

	fmt.Printf("Listening http:localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
