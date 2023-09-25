package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	RoleID   uint
	Role     Role `gorm:"foreignKey:RoleID"`
}

type Role struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func CreateUser(db *gorm.DB, username, password, email string, roleID uint) (*User, error) {
	newUser := &User{
		Username: username,
		Password: password,
		Email:    email,
		RoleID:   roleID,
	}
	if err := db.Create(newUser).Error; err != nil {
		return nil, err // Hata durumunda hatayı döndür
	}
	return newUser, nil
}

func UpdateUser(db *gorm.DB, userID uint, updateUser *User) error {
	var existingUser User
	if err := db.First(&existingUser, userID).Error; err != nil {
		return err
	}
	existingUser.Username = updateUser.Username
	existingUser.Password = updateUser.Password
	existingUser.Email = updateUser.Email
	if err := db.Save(&existingUser).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, userID uint) error {
	var existingUser User
	if err := db.First(&existingUser, userID).Error; err != nil {
		return err
	}
	if err := db.Delete(&existingUser).Error; err != nil {
		return err
	}
	return nil
}
