package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lokesh/mysql_go/pkg/models"
	"github.com/lokesh/mysql_go/pkg/utils"
)

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	book := models.GetAllBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&book)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	d := models.DeleteBook(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&d)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	var vars = mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	BookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		BookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		BookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		BookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&BookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&BookDetails)
}
