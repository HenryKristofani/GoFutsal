package models

// User represents a user account in the system
// swagger:model User
// @Description Data akun user untuk booking online
// @Description Password harus di-hash pada implementasi nyata

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role" example:"client"` // admin atau client
}
