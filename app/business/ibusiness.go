package business

import "github.com/denisacostaq/todolist-go/app/models"

type IBusiness interface {
	Validate() error
}

type ILabelBusiness interface {
	List() (labels []models.Label, err error)
	Create() (models.Label, error)
	Retrieve(id uint) (label models.Label, err error)
	Update() (models.Label, error)
	Delete(id uint) error
}

type ITaskBusiness interface {
	List() (tasks []models.Task, err error)
	Create() (models.Task, error)
	Retrieve(id uint) (task models.Task, err error)
	Update() (models.Task, error)
	Delete(id uint) error
	Priority() (uint, error)
}
