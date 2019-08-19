package modelview

import (
	"time"
	"github.com/denisacostaq/todolist-go/app/models"
	"github.com/denisacostaq/todolist-go/app/repository"
)

type LabelVm struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Color string `json:"color"`
}

type TaskVm struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
	Labels []LabelVm `json:"labels"`
}

func FromDbLabel(label models.Label) LabelVm {
	return LabelVm{
		Id: label.ID,
		Name: label.Name,
		Description: label.Description,
		Color: label.Color,
	}
}

func FromDbLabels(labels []models.Label) []LabelVm {
	var ls []LabelVm
	for _, l := range labels {
		ls = append(ls, FromDbLabel(l))
	}
	return ls
}

func FromDbTask(task models.Task) TaskVm {
	return TaskVm{
		Id: task.ID,
		Name: task.Name,
		Description: task.Description,
		DueDate: task.DueDate,
		Labels: FromDbLabels(task.Labels),
	}
}

func FromDbTasks(tasks []models.Task) []TaskVm {
	var ts []TaskVm
	for _, t := range tasks {
		ts = append(ts, FromDbTask(t))
	}
	return ts
}

func FromVmLabel(vm LabelVm, lr repository.ILabelRepository) (models.Label, error) {
	l := models.Label{}
	if vm.Id > 0 {
		if l, err := lr.Retrieve(vm.Id); err != nil {
			return l, err
		}
	}
	l.ID = vm.Id
	l.Name = vm.Name
	l.Description = vm.Description
	l.Color = vm.Color
	return l, nil
}

func FromVmLabels(mvs []LabelVm, lr repository.ILabelRepository) ([]models.Label, error) {
	var labels []models.Label
	for _, l := range mvs {
		lm, err := FromVmLabel(l, lr)
		if err != nil {
			return labels, err
		}
		labels = append(labels, lm)
	}
	return labels, nil
}

func FromVmTask(vm TaskVm, tr repository.ITaskRepository, lr repository.ILabelRepository) (models.Task, error) {
	t := models.Task{}
	if vm.Id > 0 {
		if t, err := tr.Retrieve(vm.Id); err != nil {
			return t, nil
		}
	}
	t.Name = vm.Name
	t.Description = vm.Description
	t.DueDate = vm.DueDate
	l, err := FromVmLabels(vm.Labels, lr)
	if err != nil {
		return t, err
	}
	t.Labels = l
	t.ID = vm.Id
	return t, nil
}

func FromVmTasks(tasks []TaskVm, tr repository.ITaskRepository, lr repository.ILabelRepository) ([]models.Task, error) {
	var ts []models.Task
	for _, t := range tasks {
		tm, err := FromVmTask(t, tr, lr)
		if err != nil {
			return ts, err
		}
		ts = append(ts, tm)
	}
	return ts, nil
}
