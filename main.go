package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book model or struct

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	//loop through and find with ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func createBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init router
	router := mux.NewRouter()

	//mock data
	books = append(books, Book{ID: "1", Isbn: "443323", Title: "12 Rules For Life", Author: &Author{Firstname: "Jordan", Lastname: "Peterson"}})
	books = append(books, Book{ID: "2", Isbn: "11123", Title: "Elon Musk", Author: &Author{Firstname: "Ashlee", Lastname: "Vance"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(
		http.ListenAndServe(":8000", router))

}
