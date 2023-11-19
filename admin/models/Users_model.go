package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name, Surname, Username, Email, Password, Contact string
	RoleID                                            uint
	Role                                              Role
}

func (user User) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&user)
}

func (user *User) AddRole(db *gorm.DB, newRole *Role) {
	db.Create(newRole)
	user.RoleID = newRole.ID
	user.Role = *newRole
}

func (user *User) Add(db *gorm.DB, newUser *User) {
	db.Create(newUser)
}

func (user *User) Create(db *gorm.DB) error {
	result := db.Create(user)
	return result.Error
}

func (user User) Get(where ...interface{}) User {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return user
	}
	db.Find(&user, where...)
	return user
}

func (user *User) GetAll(db *gorm.DB, where ...interface{}) []User {
	var users []User
	db.Find(&users, where...)
	return users
}

func (user *User) Update(db *gorm.DB, column string, value interface{}) {
	db.Model(user).Update(column, value)
}

func (user *User) Updates(db *gorm.DB, data User) {
	db.Model(user).Updates(data)
}

func (user *User) Delete(db *gorm.DB) {
	db.Delete(user, user.ID)
}
