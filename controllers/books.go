package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"viberate-api/database"
	"viberate-api/models"
)

func GetBooks(write http.ResponseWriter, req *http.Request) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, title, quantity FROM books")
	if err != nil {
		http.Error(write, "Unable to fetch books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		rows.Scan(&book.ID, &book.Title, &book.Quantity)
		books = append(books, book)
	}

	write.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(write).Encode(books)
	if err != nil {
		http.Error(write, "Failed to encode users to JSON", http.StatusInternalServerError)
		return
	}
}
