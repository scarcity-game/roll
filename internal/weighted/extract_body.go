package weighted

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/parse_utils"
)

func ExtractWeightedSpecification(c *gin.Context) (*Specification, error) {
	ret := &Specification{}
	if err := c.BindJSON(ret); err != nil {
		return nil, err
	} else if seed, err := parse_utils.ParseSeed(ret.StringSeed); err != nil {
		return nil, err
	} else {
		ret.Seed = seed
	}
	return ret, nil
}
