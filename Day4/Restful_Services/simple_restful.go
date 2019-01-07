package main

import (
	"encoding/json"
	"io"
	"net/http"

	"go-prog/Day4/Restful_Services/DAL"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "GETBOOKS is called...")
	//books := DAL.GetTempBooks()
	books := DAL.GetBooks()
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GetBookById is called...")
}
func PostBook(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "PostBook is called...")
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "DeleteBook is called...")
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "UpdateBook is called...")
}

func main() {
	//handle routing
	//http://localhost:8080/
	//http.HandleFunc("/", writeData)
	//using mux package
	var router = mux.NewRouter()

	//GET: http://localhost:8080/books
	router.HandleFunc("/books", GetBooks).Methods("GET")
	//GET: http://localhost:8080/books/12
	router.HandleFunc("/books/{id}", GetBookById).Methods("GET")
	//POST: http://localhost:8080/books
	router.HandleFunc("/books", PostBook).Methods("POST")
	//DELETE: http://localhost:8080/books/12
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	//PUT: http://localhost:8080/books
	router.HandleFunc("/books", UpdateBook).Methods("PUT")

	//Listen on Port 8080
	http.ListenAndServe(":8080", router)
}
