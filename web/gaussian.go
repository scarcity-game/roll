package web

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/gaussian"
	"github.com/scarcity-game/roll/web/queryparams"
	"net/http"
)

func SampleGaussian(c *gin.Context) {
	rollSpecification, err := queryparams.ExtractSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gaussianSpecification, err := gaussian.ExtractGaussianSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := rollSpecification.Roll(gaussianSpecification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
