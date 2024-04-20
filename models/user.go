package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Password string 
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}