package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	UserID    uint
	User      User
	Action    string
	Timestamp time.Time
	Details   string
}

func (log *Log) Create(db *gorm.DB) error {
	result := db.Create(log)
	return result.Error
}

func (log *Log) Get(db *gorm.DB, where ...interface{}) (Log, error) {
	var logData Log
	result := db.Where(where[0], where[1:]...).First(&logData)
	return logData, result.Error
}

func (log *Log) GetAll(db *gorm.DB, where ...interface{}) ([]Log, error) {
	var logs []Log
	result := db.Where(where[0], where[1:]...).Find(&logs)
	return logs, result.Error
}

func (log *Log) Delete(db *gorm.DB, where ...interface{}) error {
	result := db.Where(where[0], where[1:]...).Delete(log)
	return result.Error
}

func (log *Log) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&Log{})
}
