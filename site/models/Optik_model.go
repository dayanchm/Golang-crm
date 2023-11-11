package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Optik struct {
	gorm.Model
	OptikTitle, OptikSlug, OptikDescription, OptikContent, OptikPicture_url string
}

func (optik Optik) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&optik)
}

func (optik Optik) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&optik)
}

func (optik Optik) Get(where ...interface{}) Optik {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return optik
	}
	db.First(&optik, where...)
	return optik
}

func (optik Optik) GetAll(where ...interface{}) []Optik {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var optiks []Optik
	db.Find(&optiks, where...)
	return optiks
}

func (optik Optik) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&optik).Update(column, value)
}

func (optik Optik) Updates(data Optik) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&optik).Updates(data)
}

func (optik Optik) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&optik, optik.ID)
}
