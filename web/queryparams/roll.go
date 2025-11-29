package queryparams

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/roll"
	"strings"
)

func ExtractRollSpecification(c *gin.Context) (*roll.Specification, error) {
	ret := roll.DefaultSpecification()
	if value, err := collectInt("rolls", c, ret.Rolls); err != nil {
		return nil, err
	} else {
		ret.Rolls = value
	}

	if value, err := collectInt64("seed", c, ret.Seed); err != nil {
		return nil, err
	} else {
		ret.Seed = value
	}

	if value, err := collectInt("keep", c, ret.Keep); err != nil {
		return nil, err
	} else {
		ret.Keep = value
	}

	if err := collectAggregation(ret, c); err != nil {
		return nil, err
	}

	if err := collectKeepCriteria(ret, c); err != nil {
		return nil, err
	}

	return ret, nil
}

func collectKeepCriteria(specification *roll.Specification, c *gin.Context) error {
	keepCriteria := c.Query("keepCriteria")
	if keepCriteria != "" {
		switch strings.ToLower(keepCriteria) {
		case "highest":
			specification.KeepCriteria = roll.Highest
		case "lowest":
			specification.KeepCriteria = roll.Lowest
		case "middle":
			specification.KeepCriteria = roll.Middle
		default:
			return errors.New("invalid keepCriterial parameter")
		}
	}
	return nil
}
func collectAggregation(specification *roll.Specification, c *gin.Context) error {
	aggregation := c.Query("aggregation")
	if aggregation != "" {
		switch strings.ToLower(aggregation) {
		case "average":
			specification.RollAggregation = roll.Average
		case "none":
			specification.RollAggregation = roll.None
		default:
			return errors.New("invalid rollAggregation parameter")
		}
	}
	return nil
}
