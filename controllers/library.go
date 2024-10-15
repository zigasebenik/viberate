package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"viberate-api/database"
	"viberate-api/models"

	"github.com/jackc/pgx"
)

func BorrowBook(write http.ResponseWriter, req *http.Request) {

	var library models.Library

	err := json.NewDecoder(req.Body).Decode(&library)
	if err != nil {
		http.Error(write, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	err = database.DB.QueryRow(context.Background(), "SELECT id, first_name, last_name FROM users WHERE id=$1", library.UserID).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil && pgx.ErrNoRows {
		http.Error(write, "User error", http.StatusInternalServerError)
		return
	}

	var quantity int
	err = database.DB.QueryRow(context.Background(), "SELECT quantity FROM books WHERE id=$1", library.BookID).Scan(&quantity)
	if err != nil || quantity <= 0 {
		http.Error(write, "Book is not available", http.StatusInternalServerError)
		return
	}
	_, err = database.DB.Exec(context.Background(), "UPDATE books SET quantity = quantity - 1 WHERE id=$1", library.BookID)
	if err != nil {
		http.Error(write, "Unable to update book quantity", http.StatusInternalServerError)
		return
	}
	_, err = database.DB.Exec(context.Background(), "INSERT INTO library (user_id, book_id) VALUES ($1, $2)", library.UserID, library.BookID)
	if err != nil {
		http.Error(write, "Unable to record book borrowing", http.StatusInternalServerError)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode("Book borrowed")
}

func ReturnBook(write http.ResponseWriter, req *http.Request) {
	var library models.Library

	err := json.NewDecoder(req.Body).Decode(&library)
	if err != nil {
		http.Error(write, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec(context.Background(), "DELETE FROM library WHERE user_id=$1 AND book_id=$2", library.UserID, library.BookID)
	if err != nil {
		http.Error(write, "Unable to delete return record", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected() == 0 {
		http.Error(write, "No matching record found", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(context.Background(), "UPDATE books SET quantity = $1 WHERE id=$2", result.RowsAffected(), library.BookID)
	if err != nil {
		http.Error(write, "Unable to update book quantity", http.StatusInternalServerError)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode("Book returned")
}
