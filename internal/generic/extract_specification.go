package generic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/parse_utils"
	"strings"
)

func ExtractSpecification(c *gin.Context) (*Specification, error) {
	ret := DefaultSpecification()
	if value, err := parse_utils.CollectInt("rolls", c, ret.Rolls); err != nil {
		return nil, err
	} else {
		ret.Rolls = value
	}
	value := c.Query("seed")
	if seed, err := parse_utils.ParseSeed(value); err != nil {
		return nil, err
	} else {
		ret.Seed = seed
	}

	if value, err := parse_utils.CollectInt("keep", c, ret.Keep); err != nil {
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

func collectKeepCriteria(specification *Specification, c *gin.Context) error {
	keepCriteria := c.Query("keepCriteria")
	if keepCriteria != "" {
		switch strings.ToLower(keepCriteria) {
		case "highest":
			specification.KeepCriteria = Highest
		case "lowest":
			specification.KeepCriteria = Lowest
		case "middle":
			specification.KeepCriteria = Middle
		default:
			return errors.New("invalid keepCriterial parameter")
		}
	}
	return nil
}
func collectAggregation(specification *Specification, c *gin.Context) error {
	aggregation := c.Query("aggregation")
	if aggregation != "" {
		switch strings.ToLower(aggregation) {
		case "average":
			specification.RollAggregation = Average
		case "none":
			specification.RollAggregation = None
		default:
			return errors.New("invalid rollAggregation parameter")
		}
	}
	return nil
}
