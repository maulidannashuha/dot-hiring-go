package models

type User struct {
	DefaultField
	Name  string `json:"name"`
	Email string `json:"email"`
	Books []Book `json:"books" gorm:"foreignKey:UserID"`
}
