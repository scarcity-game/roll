package dice

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/generic"
	"net/http"
)

func RollDice(c *gin.Context) {
	diceString := c.Query("dice")
	advantageSpecification, err := generic.ExtractSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	diceSpecification, err := ParseDiceString(diceString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := advantageSpecification.Roll(diceSpecification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
