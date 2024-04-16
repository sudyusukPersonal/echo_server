package models

import (
	"errors"
	"time"

	"github.com/sudyusukPersonal/echo_server/database"
	"gorm.io/gorm"
)

const (
	ErrStatusServerError = "server_error"
	ErrStatusAlreadyExists = "already_exists"
	ErrStatusOK = "ok"
)

type User struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password_Hash string `json:"password_hash"`
	Created_at    time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (u *User) GerUser() User {
	return *u
}

func EmailExists(email string) (bool, error) {
	var user User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
					// レコードが見つからなかった場合、存在しない
					return false, nil
			}
			// データベースエラーの場合
			return false, err
	}
	// レコードが見つかった場合、存在する
	return true, nil
}




func NewUser(db *gorm.DB,email string) error {

	newUser := &User{
		Username: "samplename",
		Email: email,
		Password_Hash: "this is pass",
		Created_at: time.Now(),
		Updated_at: time.Now(),
		
	}
	return db.Create(newUser).Error
}