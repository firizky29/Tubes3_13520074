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
			"error": "TypeError",
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
				"error": "InvalidDNA",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "InvalidDiseaseName",
		})
	}

}
