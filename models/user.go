package models

import (
	"time"

	"gorm.io/gorm"
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


func NewUser(db *gorm.DB) error {
	
	newUser := &User{
		Username: "samplename",
		Email: "sample44@gmail.com",
		Password_Hash: "this is pass",
		Created_at: time.Now(),
		Updated_at: time.Now(),
		
	}
	return db.Create(newUser).Error
}