package controllers

import (
	"github.com/jinzhu/gorm"
)

type Controllers struct {
	DB *gorm.DB
}

func InDB(db *gorm.DB) *Controllers {
	return &Controllers{
		DB: db,
	}
}
