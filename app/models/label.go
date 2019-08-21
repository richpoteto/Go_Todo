package models

import (
	"github.com/jinzhu/gorm"
)

type Label struct {
	gorm.Model
	Name string `gorm:"type:varchar(64);not null"`
	Description string `gorm:"size:255"`
	Color string
	Priority uint `gorm:"not null"`
}
