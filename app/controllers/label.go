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

type Label struct {
	gormc.TxnController
}

// List list all labels
func (c Label) List() revel.Result {
	bl := business.NewLabel(repository.NewLabel(c.Txn), models.Label{})
	var labels []models.Label
	var err error
	if labels, err = bl.List(); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(modelview.FromDbLabels(labels))
}

// Create create a new label in database
func (c Label) Create() revel.Result {
	var mvLabel modelview.LabelVm
	if err := c.Params.BindJSON(&mvLabel); err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderError(err)
	}
	var label models.Label
	var err error
	rl := repository.NewLabel(c.Txn)
	if label, err = modelview.FromVmLabel(mvLabel, rl); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	bl := business.NewLabel(rl, label)
	if label, err = bl.Create(); err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(modelview.FromDbLabel(label))
}

// Retrieve retrieve a label
func (c Label) Retrieve(id uint) revel.Result {
	var label models.Label
	bl := business.NewLabel(repository.NewLabel(c.Txn), label)
	var err error
	if label, err = bl.Retrieve(id); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(modelview.FromDbLabel(label))
}

// Update update a label
func (c Label) Update(id uint) revel.Result {
	var vmLabel modelview.LabelVm
	if err := c.Params.BindJSON(&vmLabel); err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderError(err)
	}
	var label models.Label
	var err error
	rl := repository.NewLabel(c.Txn)
	if label, err = modelview.FromVmLabel(vmLabel, rl); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	label.ID = id
	if label, err = business.NewLabel(rl, label).Update(); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(modelview.FromDbLabel(label))
}

// Delete "delete" a label, this can be undone
func (c Label) Delete(id uint) revel.Result {
	rl := repository.NewLabel(c.Txn)
	if err := rl.Delete(id); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusOK
	return c.Render(nil)
}
