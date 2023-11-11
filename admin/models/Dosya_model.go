package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dosya struct {
	gorm.Model
	DosyaTitle, DosyaSlug, DosyaDescription, DosyaContent, Dosya_Url string
}

func (dosya Dosya) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&dosya)
}

func (dosya Dosya) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&dosya)
}

func (dosya Dosya) Get(where ...interface{}) Dosya {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return dosya
	}
	db.First(&dosya, where...)
	return dosya
}

func (dosya Dosya) GetAll(where ...interface{}) []Dosya {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var dosyas []Dosya
	db.Find(&dosyas, where...)
	return dosyas
}

func (dosya Dosya) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&dosya).Update(column, value)
}

func (dosya Dosya) Updates(data Optik) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&dosya).Updates(data)
}

func (dosya Dosya) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&dosya, dosya.ID)
}
