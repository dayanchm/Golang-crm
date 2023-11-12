package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type General struct {
	gorm.Model
	SiteTitle, FooterTitle, Logo, Favicon string
}

func (general General) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&general)
}

func (general General) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&general)
}

func (general General) Get(where ...interface{}) General {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return general
	}
	db.First(&general, where...)
	return general
}

func (general General) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&general).Update(column, value)
}

func (general General) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&general, general.ID)
}
