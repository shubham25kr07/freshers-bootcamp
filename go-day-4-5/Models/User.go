package Models

import (
	"errors"
	"go-day-4-5/Config"
	"go-day-4-5/Utils"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(u *User) (err error) {

	err = Config.DB.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func LoginCheck(username, password string, user *User) (string, error) {

	var err error
	if err = Config.DB.Where("username = ?", username).First(user).Error; err != nil {
		return "", err
	}
	match := VerifyPassword(password, user.Password)
	if match == false {
		return "", errors.New("ErrorMessage: Password  Not Matched")
	}

	token, err1 := Utils.GenerateToken(user.ID)
	if err1 != nil {
		return "", err1
	}

	return token, nil
}
