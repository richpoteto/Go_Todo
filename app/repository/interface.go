package repository

import "github.com/denisacostaq/todolist-go/app/models"

type ILabelRepository interface {
	List() (labels []models.Label, err error)
	Create(label models.Label) (models.Label, error)
	Retrieve(id uint) (label models.Label, err error)
	Update(label models.Label) (models.Label, error)
	Delete(id uint) error
}

type ITaskRepository interface {
	List() (tasks []models.Task, err error)
	Create(task models.Task) (models.Task, error)
	Retrieve(id uint) (task models.Task, err error)
	Update(task models.Task) (models.Task, error)
	Delete(id uint) error
}
