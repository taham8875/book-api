package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// create a slice of books
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	// return 404 if book not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Book{})

}

// create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	// decode the request body into a book struct
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = uuid.New().String()
	books = append(books, book)
	// return the created book
	json.NewEncoder(w).Encode(book)
}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	book.ID = params["id"]
	// decode the request body into a book struct
	json.NewDecoder(r.Body).Decode(&book)
	for index, item := range books {
		if item.ID == params["id"] {
			// update the book in the slice
			books[index] = book
			// return the updated book
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			// delete the book from the slice
			books = append(books[:index], books[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func main() {
	// initialize the slice of books
	books = append(books, Book{ID: "1", Isbn: "12345", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "67890", Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "13579", Title: "Book Three", Author: &Author{FirstName: "Jane", LastName: "Doe"}})

	// create a new router
	r := mux.NewRouter()

	// attach each path with handler (create endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// run the server
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", r)
}
