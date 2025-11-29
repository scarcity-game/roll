package weighted

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/web/output"
	"net/http"
)

func MakeWeightedChoice(c *gin.Context) {
	if weightedSpecification, err := ExtractWeightedSpecification(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err := weightedSpecification.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if result, err := weightedSpecification.Roll(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		outcome := &output.Outcome{
			Seed:        weightedSpecification.Seed,
			StringValue: result,
		}
		outcome.LogRef()
		c.IndentedJSON(http.StatusOK, outcome)
	}
}
