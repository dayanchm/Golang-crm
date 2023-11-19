package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string
}

func (role *Role) Migrate(db *gorm.DB) {
	db.AutoMigrate(role)

	roles := []Role{
		{Name: "Admin"},
		{Name: "Customer"},
		{Name: "Demo"},
		{Name: "Client"},
	}

	for _, r := range roles {
		db.FirstOrCreate(&r, Role{Name: r.Name})
	}
}
