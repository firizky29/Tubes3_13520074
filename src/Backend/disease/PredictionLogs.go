package disease

import (
	"time"

	"gorm.io/gorm"
)

/*
	All inputs should be validated first
*/

type PredictionLogs interface {
	FindAll() ([]DiseasePrediction, error)
	FindByDateComponent(year int, month int, day int) ([]DiseasePrediction, error)
	FindByDate(date time.Time) ([]DiseasePrediction, error)
	FindByDiseaseName(name string) ([]DiseasePrediction, error)
	FindByDateAndDiseaseName(name string, date time.Time) ([]DiseasePrediction, error)
	FindByNameAndDateComponent(name string, year int, month int, day int) ([]DiseasePrediction, error)
	Create(prediction DiseasePrediction) (DiseasePrediction, error)
}

type predictionLogs struct {
	db *gorm.DB
}

func NewPredictionLogs(db *gorm.DB) *predictionLogs {
	return &predictionLogs{db}
}

func (p *predictionLogs) FindAll() ([]DiseasePrediction, error) {
	var predictions []DiseasePrediction

	err := p.db.Find(&predictions).Error

	return predictions, err
}

func (p *predictionLogs) FindByDateComponent(year int, month int, day int) ([]DiseasePrediction, error) {
	var predictions []DiseasePrediction

	err := p.db.Where("year(created_at) = ? and month(created_at) = ? and day(created_at) = ?", year, month, day).Find(&predictions).Error

	return predictions, err
}

func (p *predictionLogs) FindByDate(date time.Time) ([]DiseasePrediction, error) {
	var predictions []DiseasePrediction

	err := p.db.Where("created_at = CAST(? AS DATE)", date.Format("2006-01-02 15:04:05")).Find(&predictions).Error

	return predictions, err
}

func (p *predictionLogs) FindByDiseaseName(name string) ([]DiseasePrediction, error) {
	var predictions []DiseasePrediction

	err := p.db.Where("disease_prediction = ?", name).Find(&predictions).Error

	return predictions, err
}

func (p *predictionLogs) FindByDateAndDiseaseName(name string, date time.Time) ([]DiseasePrediction, error) {
	var predictions []DiseasePrediction

	err := p.db.Where("created_at = CAST(? AS DATE) and disease_prediction = ?", date.Format("2006-01-02 15:04:05"), name).Find(&predictions).Error

	return predictions, err
}

func (p *predictionLogs) FindByNameAndDateComponent(name string, year int, month int, day int) ([]DiseasePrediction, error) {
	var predictions []DiseasePrediction

	err := p.db.Where("disease_prediction = ? and year(created_at) = ? and month(created_at) = ? and day(created_at) = ?", name, year, month, day).Find(&predictions).Error

	return predictions, err
}

func (p *predictionLogs) Create(prediction DiseasePrediction) (DiseasePrediction, error) {
	err := p.db.Create(&prediction).Error

	return prediction, err
}
