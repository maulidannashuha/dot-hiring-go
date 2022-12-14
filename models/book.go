package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}
