package uniform

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/generic"
	"net/http"
)

func SampleUniform(c *gin.Context) {
	advantageSpecification, err := generic.ExtractSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uniformSpecification, err := ExtractUniformSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := advantageSpecification.Roll(uniformSpecification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
