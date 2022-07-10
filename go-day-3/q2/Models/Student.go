package Models

import (
	"fmt"
	"q2/Config"
)

func GetAllUsers(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}
func GetUserByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(student *Student) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}
