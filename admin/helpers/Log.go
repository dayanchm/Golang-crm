package helpers

import (
	"blog/admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"fmt"
)

func LogAction(userID uint, action string, details string) {
	db, err := gorm.Open(mysql.Open(models.Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	log := models.Log{
		UserID:    userID,
		Action:    action,
		Timestamp: time.Now(),
		Details:   details,
	}
	log.Create(db)
}
