package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name string `gorm:"not null"`
	Description string `gorm:"size:255"`
	DueDate time.Time `gorm:"not null"`
	Labels []Label `gorm:"many2many:task_labels"`
	Priority float32 `gorm:"-"` // do not persist this field
}
