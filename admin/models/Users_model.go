package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func (user User) CreateUser() {
	Database.Db.Create(&user)
}
func (user User) GetAllUsers(where ...interface{}) []User {
	var users []User
	Database.Db.Find(&users, where...)
	return users
}
func (user User) GetSingleUser(where ...interface{}) User {
	Database.Db.First(&user, where...)
	return user
}
func UpdateUser(column string, value interface{}) {
	var user User
	Database.Db.First(&user, 1)
	Database.Db.Model(&user).Update(column, value)
}
func (user User) UpdatesUser(data User) {
	Database.Db.First(&user, 1)
	Database.Db.Model(&user).Updates(data)
}
func (user User) DeleteUser() {
	Database.Db.Delete(&user, user.ID)
}
