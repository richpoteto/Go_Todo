package controllers

import (
	"github.com/denisacostaq/todolist-go/app/business"
	"github.com/denisacostaq/todolist-go/app/controllers/modelview"
	"github.com/denisacostaq/todolist-go/app/models"
	"github.com/denisacostaq/todolist-go/app/repository"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"

	"github.com/revel/revel"
	"net/http"
)

type Task struct {
	gormc.TxnController
}

// List list all tasks
func (c Task) List() revel.Result {
	bt := business.NewTask(repository.NewTask(c.Txn), models.Task{})
	var tasks []models.Task
	var err error
	if tasks, err = bt.List(); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(modelview.FromDbTasks(tasks))
}

// Create create a new task in database
func (c Task) Create() revel.Result {
	var vmTask modelview.TaskVm
	if err := c.Params.BindJSON(&vmTask); err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderError(err)
	}
	var err error
	var task models.Task
	tr := repository.NewTask(c.Txn)
	if task, err = modelview.FromVmTask(vmTask, tr, repository.NewLabel(c.Txn)); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	bt := business.NewTask(tr, task)
	if task, err = bt.Create(); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(modelview.FromDbTask(task))
}

// Retrieve retrieve a contact
func (c Task) Retrieve(id uint) revel.Result {
	var task models.Task
	bt := business.NewTask(repository.NewTask(c.Txn), models.Task{})
	var err error
	if task, err = bt.Retrieve(id); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(modelview.FromDbTask(task))
}

// Update update a task
func (c Task) Update(id uint) revel.Result {
	var vmTask modelview.TaskVm
	if err := c.Params.BindJSON(&vmTask); err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderError(err)
	}
	tr := repository.NewTask(c.Txn)
	var err error
	var task models.Task
	if task, err = modelview.FromVmTask(vmTask, tr, repository.NewLabel(c.Txn)); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	task.ID = id
	bt := business.NewTask(tr, task)
	if task, err = bt.Update(); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(modelview.FromDbTask(task))
}

// Delete "delete" a task, this can be undone
func (c Task) Delete(id uint) revel.Result {
	bt := business.NewTask(repository.NewTask(c.Txn), models.Task{})
	if err := bt.Delete(id); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(nil)
}
