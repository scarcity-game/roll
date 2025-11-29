package gaussian

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/parse_utils"
)

func ExtractGaussianSpecification(c *gin.Context) (*Specification, error) {
	ret := DefaultSpecification()
	if value, err := parse_utils.CollectFloat64(parse_utils.MeanParam, c, ret.Mean); err != nil {
		return nil, err
	} else {
		ret.Mean = value
	}

	if value, err := parse_utils.CollectFloat64(parse_utils.StddevParam, c, ret.Stddev); err != nil {
		return nil, err
	} else {
		ret.Stddev = value
	}

	if value, err := parse_utils.CollectFloat64(parse_utils.MaxParam, c, ret.Max); err != nil {
		return nil, err
	} else {
		ret.Max = value
	}
	if value, err := parse_utils.CollectFloat64(parse_utils.MinParam, c, ret.Min); err != nil {
		return nil, err
	} else {
		ret.Min = value
	}
	return ret, nil
}
