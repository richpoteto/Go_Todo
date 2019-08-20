package business

import (
	"errors"
	"github.com/denisacostaq/todolist-go/app/models"
	"github.com/denisacostaq/todolist-go/app/repository"
	"time"
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
	if time.Now().After(t.tm.DueDate) {
		return errors.New("you can not refer to the pass in the due date")
	}
	return nil
}

func (t Task) List() ([]models.Task, error) {
	tms, err := t.tr.List()
	if err != nil {
		return []models.Task{}, err
	}
	for idxT := range tms {
		task := Task{t.tr, tms[idxT]}
		if tms[idxT].Priority, err = task.Priority(); err != nil {
			return []models.Task{}, err
		}
	}
	return tms, err
}

func (t Task) Create() (models.Task, error) {
	if err := t.Validate(); err != nil {
		return models.Task{}, err
	}
	var err error
	t.tm, err = t.tr.Create(t.tm)
	if t.tm.Priority, err = t.Priority(); err != nil {
		return models.Task{}, err
	}
	return t.tm, err
}

func (t Task) Retrieve(id uint) (models.Task, error) {
	var err error
	t.tm, err = t.tr.Retrieve(id)
	if t.tm.Priority, err = t.Priority(); err != nil {
		return models.Task{}, err
	}
	return t.tm, err
}

func (t Task) Update() (models.Task, error) {
	if err := t.Validate(); err != nil {
		return models.Task{}, err
	}
	var err error
	t.tm, err = t.tr.Update(t.tm)
	if t.tm.Priority, err = t.Priority(); err != nil {
		return models.Task{}, err
	}
	return t.tm, err
}

func (t Task) Delete(id uint) error {
	return t.tr.Delete(id)
}

func (t Task) Priority() (float32, error) {
	var priority uint
	priority = 0
	var tdb models.Task
	var err error
	if tdb, err = t.tr.Retrieve(t.tm.ID); err != nil {
		return 0, err
	}
	for idxT := range tdb.Labels {
		priority += tdb.Labels[idxT].Priority
	}
	if len(tdb.Labels) == 0 {
		return 0, nil
	}
	return float32(priority)/float32(len(tdb.Labels)), nil
}
