package business

import (
	"errors"
	"github.com/denisacostaq/todolist-go/app/models"
	"github.com/denisacostaq/todolist-go/app/repository"
)

type Task struct {
	tr repository.ITaskRepository
	tm models.Task
}

func NewTask(tr repository.ITaskRepository, tm models.Task) *Task {
	return &Task{tr: tr, tm: tm}
}

func (t Task) Validate() error {
	if len(t.tm.Name) == 0 {
		return errors.New("name is required in Task")
	}
	return nil
}

func (t Task) List() ([]models.Task, error) {
	return t.tr.List()
}

func (t Task) Create() (models.Task, error) {
	if err := t.Validate(); err != nil {
		return models.Task{}, err
	}
	var err error
	t.tm, err = t.tr.Create(t.tm)
	return t.tm, err
}

func (t Task) Retrieve(id uint) (models.Task, error) {
	var err error
	t.tm, err = t.tr.Retrieve(id)
	return t.tm, err
}

func (t Task) Update() (models.Task, error) {
	if err := t.Validate(); err != nil {
		return models.Task{}, err
	}
	var err error
	t.tm, err = t.tr.Update(t.tm)
	return t.tm, err
}

func (t Task) Delete(id uint) error {
	return t.tr.Delete(id)
}
