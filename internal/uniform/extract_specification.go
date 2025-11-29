package uniform

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/parse_utils"
)

func ExtractUniformSpecification(c *gin.Context) (*Specification, error) {
	ret := DefaultSpecification()
	if m, err := parse_utils.CollectFloat64(parse_utils.MaxParam, c, ret.Max); err != nil {
		return nil, err
	} else {
		ret.Max = m
	}
	if m, err := parse_utils.CollectFloat64(parse_utils.MinParam, c, ret.Min); err != nil {
		return nil, err
	} else {
		ret.Min = m
	}
	return ret, nil
}
