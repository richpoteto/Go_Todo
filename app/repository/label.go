package repository

import (
	"errors"
	"github.com/denisacostaq/todolist-go/app/models"
	"github.com/jinzhu/gorm"
)

type Label struct {
	tx *gorm.DB
}

func NewLabel(tx *gorm.DB) ILabelRepository {
	return &Label{tx: tx}
}

func (l Label) List() (labels []models.Label, err error) {
	if res := l.tx.Find(&labels); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		//	log.Panic("Error retrieving from database: ", res)
		return labels, errors.New("error retrieving from database")
	}
	return labels, err
}

func (l Label) Create(label models.Label) (models.Label, error) {
	if res := l.tx.Create(&label); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		//log.Panic("Error creating in database: ", res)
		return models.Label{}, errors.New("error creating in database")
	}
	return label, nil
}

func (l Label) Retrieve(id uint) (label models.Label, err error) {
	if res := l.tx.First(&label, id); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		//log.Panic("Error retrieving from database: ", res)
		return label, errors.New("error retrieving from database")
	}
	return label, err
}

func (l Label) Update(label models.Label) (models.Label, error) {
	if res := l.tx.Save(&label); len(res.GetErrors()) != 0 {
		// TODO(denisacostaq@gmail.com): add logging
		// log.Panic("Error updating label with id %u: ", c.ID, res)
		return label, errors.New("error updating Label")
	}
	return label, nil
}

func (l Label) Delete(id uint) error {
	var label models.Label
	label.ID = id
	if res := l.tx.Delete(&label); len(res.GetErrors()) != 0 {
		// log.Panic("Error deleting Label with id %u: ", id, res)
		return errors.New("error deleting Label")
	}
	return nil
}
