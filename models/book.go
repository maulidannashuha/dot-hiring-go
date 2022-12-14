package models

type Book struct {
	DefaultField
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}
