package repository

import "projects/web-chat-app/database"

type User struct {
	ID       int64
	Username string
	Password string
}

func (user *User) InsertToDB() (User, error) {
	DB := database.GetDB()
	_, err := DB.Exec(`INSERT INTO "USER" (username, password) VALUES($1, $2)`, 
						user.Username, user.Password)

	if err != nil {
		return User{}, err
	}

	return *user, nil
}