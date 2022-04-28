package service

import (
	"backend/algorithm"
	"backend/disease"
	"backend/transaction"
	"fmt"

	"github.com/shopspring/decimal"
)

type PredictionService interface {
	FindAll() ([]disease.DiseasePrediction, error)
	FindLogs(name string, date string) ([]transaction.PredictionResponse, error)
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

func (p *predictionService) FindLogs(name string, date string) ([]transaction.PredictionResponse, error) {
	var Logs []transaction.PredictionResponse
	var curPredictions []disease.DiseasePrediction
	var err error
	if name != "" && date != "" {
		curPredictions, err = p.predictionLogs.FindByDateAndDiseaseName(date, name)
	} else if name != "" {
		curPredictions, err = p.predictionLogs.FindByDiseaseName(name)
	} else if date != "" {
		curPredictions, err = p.predictionLogs.FindByDate(date)
	}

	if err != nil {
		return []transaction.PredictionResponse{}, err
	}
	for _, prediction := range curPredictions {
		y, m, d := prediction.CreatedAt.Date()
		similarityResult := prediction.SimilarityLevel.String() + "%"
		result := fmt.Sprintf("%d %s %d - %s - %s - %s - %s", d, m, y, prediction.UserName, prediction.DiseasePrediction, similarityResult, prediction.Status)
		Log := transaction.PredictionResponse{Result: result}
		Logs = append(Logs, Log)
	}
	return Logs, nil
}

func (p *predictionService) Create(prediction transaction.PredictionRequest, DNADisease string) (transaction.PredictionResponse, error) {
	var status string
	if prediction.Algorithm == "Boyer Moore" {
		algorithm.BoyerMooreSearch(prediction.DNA, DNADisease, &status)
	} else {
		algorithm.KMPSearch(prediction.DNA, DNADisease, &status)
	}
	similarity := algorithm.SimilarityLevel(prediction.DNA, DNADisease)
	if status == "false" && similarity.Cmp(decimal.NewFromFloat(80.0)) == 1 {
		status = "true"
	}
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
	similarityResult := CreatedLogs.SimilarityLevel.String() + "%"

	PredictionResult := fmt.Sprintf("%d %s %d - %s - %s - %s - %s", d, m, y, CreatedLogs.UserName, CreatedLogs.DiseasePrediction, similarityResult, CreatedLogs.Status)

	return transaction.PredictionResponse{Result: PredictionResult}, nil
}
