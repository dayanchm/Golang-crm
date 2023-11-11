package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Optikokuyucu struct {
	gorm.Model
	OptikTitle, OptikSlug, OptikDescription, OptikContent, OptikPicture_url string
}

func (optik Optikokuyucu) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&optik)
}

func (optik Optikokuyucu) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&optik)
}

func (optik Optikokuyucu) Get(where ...interface{}) Optikokuyucu {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return optik
	}
	db.First(&optik, where...)
	return optik
}

func (optik Optikokuyucu) GetAll(where ...interface{}) []Optikokuyucu {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var optiks []Optikokuyucu
	db.Find(&optiks, where...)
	return optiks
}

func (optik Optikokuyucu) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&optik).Update(column, value)
}

func (optik Optikokuyucu) Updates(data Optikokuyucu) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&optik).Updates(data)
}

func (optik Optikokuyucu) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&optik, optik.ID)
}
