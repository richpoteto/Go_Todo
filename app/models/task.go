package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name string
	DueDate time.Time
}
