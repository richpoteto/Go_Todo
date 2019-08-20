package repository

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/denisacostaq/todolist-go/app/models"
	"github.com/jinzhu/gorm"
)

type Task struct {
	tx *gorm.DB
}

func NewTask(tx *gorm.DB) ITaskRepository {
	return &Task{tx: tx}
}

func (t Task) List() (tasks []models.Task, err error) {
	if res := t.tx.Preload("Labels").Find(&tasks); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		//	log.Panic("Error retrieving from database: ", res)
		return tasks, errors.New("error retrieving from database")
	}
	return tasks, nil
}

func (t Task) Create(task models.Task) (models.Task, error) {
	spew.Dump("tasksss", task)
	if res := t.tx.Create(&task); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		//log.Panic("Error creating in database: ", res)
		return task, errors.New("error creating in database")
	}
	return task, nil
}

func (t Task) Retrieve(id uint) (task models.Task, err error) {
	if res := t.tx.Preload("Labels").First(&task, id); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		//log.Panic("Error retrieving from database: ", res)
		return task, errors.New("error retrieving from database")
	}
	return task, err
}

func (t Task) Update(task models.Task) (models.Task, error) {
	t.tx.Model(&task).Association("Labels").Replace(task.Labels)
	tannedDb := t.tx.Set("gorm:save_associations", true)
	tannedDb = tannedDb.Set("gorm:association_save_reference", true)
	tannedDb = tannedDb.Set("gorm:association_autoupdate", false)
	tannedDb = tannedDb.Set("gorm:association_autocreate", false)
	if res := tannedDb.Save(&task); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		// log.Panic("Error updating task with id %u: ", c.ID, res)
		return models.Task{}, errors.New("error updating Task")
	}
	return task, nil
}

func (t Task) Delete(id uint) error {
	var task models.Task
	task.ID = id
	if res := t.tx.Delete(&task); len(res.GetErrors()) != 0 {
		// log.Panic("Error deleting Task with id %u: ", id, res)
		return errors.New("error deleting Task")
	}
	return nil
}
