package models

type User struct {
	IDuser   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
