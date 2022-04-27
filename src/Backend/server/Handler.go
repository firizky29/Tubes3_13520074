package server

import (
	"backend/algorithm"
	"backend/transaction"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) PostDisease(c *gin.Context) {
	var disease transaction.DiseaseRequest

	err := c.ShouldBindJSON(&disease)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while trying to read input, please make sure your input is valid",
		})
	} else if algorithm.ValidateName(disease.DiseaseName) {
		if algorithm.ValidateDNA(disease.DNA) {
			temp, _ := s.DiseaseService.FindByName(disease.DiseaseName)
			if temp.Name == "" {
				_, err := s.DiseaseService.Create(disease)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "Error while trying to insert, please try again",
					})
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("Disease with name %s already exists", disease.DiseaseName),
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid DNA input, it should only contains A, T, C, or G character.",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Name Input.",
		})
	}
}

func (s *Server) PostPrediction(c *gin.Context) {
	var prediction transaction.PredictionRequest

	err := c.ShouldBindJSON(&prediction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while trying to read input, please make sure your input is valid.",
		})
	} else if algorithm.ValidateName(prediction.UserName) {
		if algorithm.ValidateDNA(prediction.DNA) {
			if algorithm.ValidateAlgorithm(prediction.Algorithm) {
				temp, _ := s.DiseaseService.FindByName(prediction.DiseasePrediction)
				if temp.Name != "" {
					PredictionResult, err := s.PredictionService.Create(prediction, temp.DNA)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"error": "Error while trying to retrieve data, please try again.",
						})
					} else {
						c.JSON(http.StatusOK, gin.H{
							"Result": PredictionResult.Result,
						})
					}
				} else {
					c.JSON(http.StatusOK, gin.H{
						"error": "Cannot find the specified disease.",
					})
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid Algorithm.",
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid DNA input, it should only contains A, T, C, or G character.",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Name Input.",
		})
	}
}
