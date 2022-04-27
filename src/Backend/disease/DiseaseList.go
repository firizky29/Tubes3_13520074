package disease

import "gorm.io/gorm"

/*
	All inputs should be validated first
*/

type DiseaseList interface {
	FindAll() ([]Disease, error)
	FindByName(name string) (Disease, error)
	Create(disease Disease) (Disease, error)
}

type diseaseList struct {
	db *gorm.DB
}

func NewDisease(db *gorm.DB) *diseaseList {
	return &diseaseList{db}
}

func (d *diseaseList) FindAll() ([]Disease, error) {
	var diseases []Disease

	err := d.db.Find(&diseases).Error

	return diseases, err
}

func (d *diseaseList) FindByName(name string) (Disease, error) {
	var disease Disease

	err := d.db.Where("name = ?", name).Find(&disease).Error

	return disease, err
}

func (d *diseaseList) Create(disease Disease) (Disease, error) {
	err := d.db.Create(&disease).Error

	return disease, err
}
