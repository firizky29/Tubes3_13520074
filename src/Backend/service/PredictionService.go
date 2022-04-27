package service

import (
	"backend/disease"
	"backend/transaction"
)

type PredictionService interface {
	FindAll() ([]disease.DiseasePrediction, error)
	FindLogs(prediction transaction.LogsRequest) ([]transaction.PredictionResponse, error)
	ProcessStringMatching(prediction transaction.PredictionRequest) (transaction.PredictionResponse, error)
	Create(prediction transaction.PredictionRequest) (disease.DiseasePrediction, error)
}

type predictionService struct {
	predictionLogs disease.PredictionLogs
}

func NewPredictionService(predictionLogs disease.PredictionLogs) *predictionService {
	return &predictionService{predictionLogs}
}

func (p *predictionService) FindAll() ([]disease.DiseasePrediction, error) {
	return p.predictionLogs.FindAll()
}

func (p *predictionService) FindLogs(prediction transaction.LogsRequest) ([]transaction.PredictionResponse, error) {
	return nil, nil
}

func (p *predictionService) ProcessStringMatching(prediction transaction.PredictionRequest) (transaction.PredictionResponse, error) {
	return transaction.PredictionResponse{}, nil
}

func (p *predictionService) Create(prediction transaction.PredictionRequest) (disease.DiseasePrediction, error) {
	return disease.DiseasePrediction{}, nil
}
