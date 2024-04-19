package repositories

import (
	"database/sql"
	"log"
	"projects/web-chat-app/database"
)

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
 
func GetUserFromDB(username string) (*User, error) {
	DB := database.GetDB()
	result := DB.QueryRow(`SELECT id,username FROM "USER" WHERE username = $1`, username)

	var user User
	if err := result.Scan(&user.ID, &user.Username); err != nil{
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error get data : ", err)
		return nil, err
	}

	return &user, nil
}