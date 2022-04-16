package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/runan13/gocoin/blockchain"
	"github.com/runan13/gocoin/utils"
	"log"
	"net/http"
	"strconv"
)

var port string

type url string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func handleDoc(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See a Block",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func handleBlocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func handleBlock(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(height)
	if err == blockchain.ErrorNotFound {
		json.NewEncoder(rw).Encode(errorResponse{fmt.Sprint(err)})
	} else {
		json.NewEncoder(rw).Encode(block)
	}
}

func Start(aPort int) {
	r := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	r.HandleFunc("/", handleDoc).Methods("GET")
	r.HandleFunc("/blocks", handleBlocks).Methods("GET", "POST")
	r.HandleFunc("/block/{height:[0-9]+}", handleBlock).Methods("GET")

	fmt.Printf("Listening http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
