package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func handleDoc(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

}

func main() {
	http.HandleFunc("/", handleDoc)

	fmt.Printf("Listening http:localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
