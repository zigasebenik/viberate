package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"viberate-api/database"
	"viberate-api/models"
)

func CreateUser(write http.ResponseWriter, req *http.Request) {
	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(write, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(context.Background(), "INSERT INTO users (first_name, last_name) VALUES ($1, $2)", user.FirstName, user.LastName)
	if err != nil {
		http.Error(write, "Unable to insert user", http.StatusInternalServerError)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode("User " + user.FirstName + ", " + user.LastName + " added")
}

func GetUsers(write http.ResponseWriter, req *http.Request) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, first_name, last_name FROM users")
	if err != nil {
		http.Error(write, "Unable to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			fmt.Println("Failed to scan row: ", err)
		}

		users = append(users, user)
	}

	write.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(write).Encode(users)
	if err != nil {
		http.Error(write, "Failed to encode users to JSON", http.StatusInternalServerError)
		return
	}
}
