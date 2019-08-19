package business

import (
	"errors"
	"github.com/denisacostaq/todolist-go/app/models"
	"github.com/denisacostaq/todolist-go/app/repository"
)

type Label struct {
	lr repository.ILabelRepository
	lm models.Label
}

func NewLabel(lr repository.ILabelRepository, lm models.Label) *Label {
	return &Label{lr: lr, lm: lm}
}

func (l Label) Validate() error {
	if len(l.lm.Name) == 0 {
		return errors.New("name is required in Label")
	}
	return nil
}

func (l Label) List() ([]models.Label, error) {
	return l.lr.List()
}

func (l Label) Create() (models.Label, error) {
	if err := l.Validate(); err != nil {
		return models.Label{}, err
	}
	var err error
	l.lm, err = l.lr.Create(l.lm)
	return l.lm, err
}

func (l Label) Retrieve(id uint) (models.Label, error) {
	var err error
	l.lm, err = l.lr.Retrieve(id)
	return l.lm, err
}

func (l Label) Update() (models.Label, error) {
	if err := l.Validate(); err != nil {
		return models.Label{}, err
	}
	var err error
	l.lm, err = l.lr.Update(l.lm)
	return l.lm, err
}

func (l Label) Delete(id uint) error {
	return l.lr.Delete(id)
}
