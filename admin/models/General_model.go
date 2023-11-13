package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type General_Setting struct {
	gorm.Model
	SiteTitle, FooterTitle, Logo, DarkLogo, Favicon string
}

func (general General_Setting) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&general)
}

func (general General_Setting) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&general)
}

func (general General_Setting) Get(where ...interface{}) General_Setting {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return general
	}
	db.First(&general, where...)
	return general
}

func (general General_Setting) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&general).Update(column, value)
}

func (general General_Setting) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&general, general.ID)
}
