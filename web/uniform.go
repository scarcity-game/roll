package web

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/web/queryparams"
	"net/http"
)

func SampleUniform(c *gin.Context) {
	rollSpecification, err := queryparams.ExtractRollSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uniformSpecification, err := queryparams.ExtractUniformSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := rollSpecification.Roll(uniformSpecification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
