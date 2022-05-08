package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"math/rand"
	//  "strconv"
	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Init books var as a slice Book struct

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //mock id - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(10000000)) //mock id - not safe
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data  -@todo -implement DB

	books = append(books, Book{ID: "1", Isbn: "1234565", Title: "Java 入门", Author: &Author{FirstName: "James ", LastName: "Gosling"}})
	books = append(books, Book{ID: "2", Isbn: "7654344", Title: "C# 入门", Author: &Author{FirstName: "Anders", LastName: "Hejlsberg"}})
	//Router Handlers /EndPoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")           // curl http://localhost:8080/api/books
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")       //>curl http://localhost:8080/api/books/1
	r.HandleFunc("/api/books", createBook).Methods("POST")        // curl  -d  "{\"isbn\":\"35444\",\"title\":\"C# over\",\"author\":{\"firstname\":\"victor\",\"lastname\":\"test\"}}" -H "Content-Type:application/json" -H "Accept:application/json"  http://localhost:8080/api/books -X POST
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")    //-d  "{\"isbn\":\"99999\",\"title\":\"test\",\"author\":{\"firstname\":\"victor\",\"lastname\":\"test\"}}" -H "Content-Type:application/json" -H "Accept:application/json"  http://localhost:8080/api/books/1 -X PUT
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE") // curl http://localhost:8080/api/books/1 -X DELETE
	log.Fatal(http.ListenAndServe(":8080", r))
}
