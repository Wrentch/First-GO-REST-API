package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Atricle struct {
	Title    string `json:"Title"`
	Desc     string `json:"Description"`
	Contrent string `json:"Content"`
}

type Articles []Atricle

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Atricle{Title: "Test title", Desc: "Test description", Contrent: "ehhee content"},
	}

	fmt.Println("End point hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func TestPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testing post method hit")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/art", allArticles).Methods("GET")
	myRouter.HandleFunc("/art", TestPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequest()
}
