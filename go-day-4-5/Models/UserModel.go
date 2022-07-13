package Models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) TableName() string {
	return "user"
}
func (u *User) BeforeSave() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = strings.TrimSpace(u.Username)

	return nil
}
