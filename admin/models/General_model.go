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
func (general *General_Setting) GetLatestGeneralSetting() (General_Setting, error) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		return *general, fmt.Errorf("Veritabanına bağlanırken bir hata oluştu: %v", err)
	}

	var setting General_Setting
	result := db.Order("created_at desc").First(&setting)

	if result.Error != nil {
		return *general, fmt.Errorf("Veritabanından veri çekme işlemi başarısız: %v", result.Error)
	}

	return setting, nil
}
func (general General_Setting) Add() error {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Veritabanına bağlanırken bir hata oluştu: %v", err)
	}

	result := db.Create(&general)
	if result.Error != nil {
		return fmt.Errorf("Veritabanına ekleme işlemi başarısız: %v", result.Error)
	}

	return nil
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

func (general General_Setting) GetAll(where ...interface{}) []General_Setting {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var generals []General_Setting
	db.Find(&generals, where...)
	return generals
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

func GetLatestGeneralSetting() (General_Setting, error) {
	var setting General_Setting
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return setting, err
	}
	db.Order("created_at desc").First(&setting)
	return setting, nil
}
