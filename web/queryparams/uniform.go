package queryparams

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/uniform"
)

func ExtractUniformSpecification(c *gin.Context) (*uniform.Specification, error) {
	ret := uniform.DefaultSpecification()
	if m, err := collectFloat64(maxParam, c, ret.Max); err != nil {
		return nil, err
	} else {
		ret.Max = m
	}
	if m, err := collectFloat64(minParam, c, ret.Min); err != nil {
		return nil, err
	} else {
		ret.Min = m
	}
	return ret, nil
}
