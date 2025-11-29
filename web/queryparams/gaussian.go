package queryparams

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/gaussian"
)

func ExtractGaussianSpecification(c *gin.Context) (*gaussian.Specification, error) {
	ret := gaussian.DefaultSpecification()
	if value, err := collectFloat64(meanParam, c, ret.Mean); err != nil {
		return nil, err
	} else {
		ret.Mean = value
	}

	if value, err := collectFloat64(stddevParam, c, ret.Stddev); err != nil {
		return nil, err
	} else {
		ret.Stddev = value
	}

	if value, err := collectFloat64(maxParam, c, ret.Max); err != nil {
		return nil, err
	} else {
		ret.Max = value
	}
	if value, err := collectFloat64(minParam, c, ret.Min); err != nil {
		return nil, err
	} else {
		ret.Min = value
	}
	return ret, nil
}
