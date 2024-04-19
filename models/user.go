package models

type User struct {
	ID       int64
	Username string
	Password string
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}