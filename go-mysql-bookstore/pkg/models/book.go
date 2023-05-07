package models

import (
	"go-mysql-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}
