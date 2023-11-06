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

	// Önce, role_id'nin geçerli bir Role'a ait olduğunu doğrula
	var role model.Role
	if err := uc.DB.First(&role, userInput.RoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz role_id"})
		return
	}

	// Veritabanı bağlantısı kontrolü
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

func (uc *UserController) UpdateUser(c *gin.Context) {
	var userInput model.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := userInput.ID
	var user model.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}
	if err := uc.DB.Model(&user).Updates(userInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
func (uc *UserController) GetOneUser(c *gin.Context) {
	userID := c.Param("id")

	var user model.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	var users []model.User
	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := uc.DB.Delete(&model.User{}, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı başarıyla silindi"})
}
