package controller

import (
	"crm/model"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB    *gorm.DB
	SQLDB *sql.DB
}

func NewUserControllerWithDB(db *sql.DB, gormDB *gorm.DB) *UserController {
	return &UserController{SQLDB: db, DB: gormDB}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userInput model.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if uc.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Veritabanı bağlantısı eksik"})
		return
	}

	result := uc.DB.Create(&userInput)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, userInput)
}
