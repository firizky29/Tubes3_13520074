package service

import (
	"backend/disease"
	"backend/transaction"
)

type DiseaseService interface {
	FindAll() ([]disease.Disease, error)
	Create(diseaseRequest transaction.DiseaseRequest) (disease.Disease, error)
}

type diseaseService struct {
	diseaseList disease.DiseaseList
}

func NewDiseaseService(diseaseList disease.DiseaseList) *diseaseService {
	return &diseaseService{diseaseList}
}

func (d *diseaseService) FindAll() ([]disease.Disease, error) {
	return d.diseaseList.FindAll()
}

func (d *diseaseService) Create(diseaseRequest transaction.DiseaseRequest) (disease.Disease, error) {
	newDisease := disease.Disease{
		Name: diseaseRequest.DiseaseName,
		DNA:  diseaseRequest.DNA,
	}
	return d.diseaseList.Create(newDisease)
}
