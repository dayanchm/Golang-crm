package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Slider struct {
	gorm.Model
	Slider_Title, Slider_Slug, Slider_Picture_url string
	CategoryID                                    int
}

func (slider Slider) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&slider)
}

func (slider Slider) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&slider)
}

func (slider Slider) Get(where ...interface{}) Slider {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return slider
	}
	db.First(&slider, where...)
	return slider
}

func (slider Slider) GetAll(where ...interface{}) []Slider {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var sliders []Slider
	db.Find(&sliders, where...)
	return sliders
}

func (slider Slider) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&slider).Update(column, value)
}

func (slider Slider) Updates(data Slider) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&slider).Updates(data)
}

func (slider Slider) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&slider, slider.ID)
}
