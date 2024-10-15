package models

type Book struct {
	ID       int    `json:"id"`
	Quantity int    `json:"quantity"`
	Title    string `json:"title"`
}
