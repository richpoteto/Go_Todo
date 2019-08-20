package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name string
	Description string
	DueDate time.Time
	Labels []Label `gorm:"many2many:task_labels"`
	Priority float32 `gorm:"-"` // do not persist this field
}
