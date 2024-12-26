package models

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Username   string `json:"username" gorm:"unique"`
	Phone      int    `json:"phone" gorm:"unique"`
	Department string `json:"department"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}
