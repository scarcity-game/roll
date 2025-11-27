package web

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/dice"
	"github.com/scarcity-game/roll/web/queryparams"
	"net/http"
)

func RollDice(c *gin.Context) {
	diceString := c.DefaultQuery("dice", "1d6")
	rollSpecification, err := queryparams.ExtractSpecification(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	diceSpecification, err := dice.ParseDiceString(diceString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := rollSpecification.Roll(diceSpecification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
