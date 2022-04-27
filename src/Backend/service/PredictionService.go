package service

import (
	"backend/algorithm"
	"backend/disease"
	"backend/transaction"
	"fmt"
)

type PredictionService interface {
	FindAll() ([]disease.DiseasePrediction, error)
	FindLogs(prediction transaction.LogsRequest) ([]transaction.PredictionResponse, error)
	Create(prediction transaction.PredictionRequest, DNADisease string) (transaction.PredictionResponse, error)
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

func (p *predictionService) Create(prediction transaction.PredictionRequest, DNADisease string) (transaction.PredictionResponse, error) {
	var status string
	if prediction.Algorithm == "Boyer Moore" {
		algorithm.BoyerMooreSearch(prediction.DNA, DNADisease, &status)
	} else {
		algorithm.KMPSearch(prediction.DNA, DNADisease, &status)
	}
	similarity := algorithm.SimilarityLevel(prediction.DNA, DNADisease)
	NewPrediction := disease.DiseasePrediction{
		UserName:          prediction.UserName,
		DNA:               prediction.DNA,
		DiseasePrediction: prediction.DiseasePrediction,
		SimilarityLevel:   similarity,
		Status:            status,
	}

	CreatedLogs, err := p.predictionLogs.Create(NewPrediction)
	if err != nil {
		return transaction.PredictionResponse{}, err
	}

	y, m, d := CreatedLogs.CreatedAt.Date()
	similarityResult := CreatedLogs.SimilarityLevel.String()

	PredictionResult := fmt.Sprintf("%d %s %d - %s - %s - %s - %s", d, m, y, CreatedLogs.UserName, CreatedLogs.DiseasePrediction, similarityResult, CreatedLogs.Status)

	return transaction.PredictionResponse{Result: PredictionResult}, nil
}
