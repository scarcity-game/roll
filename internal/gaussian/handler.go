package gaussian

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/generic"
	"net/http"
)

func SampleGaussian(c *gin.Context) {
	advantageSpecification, err := generic.ExtractSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gaussianSpecification, err := ExtractGaussianSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := advantageSpecification.Roll(gaussianSpecification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
