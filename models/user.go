package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"` // super_admin, admin_level_2, admin_level_3, user
	// ...existing code...
}
