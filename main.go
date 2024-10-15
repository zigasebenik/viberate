package main

import (
	"fmt"
	"log"
	"net/http"
	"viberate-api/controllers"
	"viberate-api/database"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB("postgres://postgres:root@localhost:5432/viberate?sslmode=disable")

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", controllers.GetUsers).Methods("GET")    //OK
	api.HandleFunc("/users", controllers.CreateUser).Methods("POST") //OK
	api.HandleFunc("/books", controllers.GetBooks).Methods("GET")    //OK
	api.HandleFunc("/borrow", controllers.BorrowBook).Methods("POST")
	api.HandleFunc("/return", controllers.ReturnBook).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	defer database.CloseDB()
}
