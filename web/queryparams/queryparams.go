package queryparams

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/roll"
	"strconv"
	"strings"
)

func ExtractSpecification(c *gin.Context) (*roll.Specification, error) {
	ret := roll.DefaultSpecification()
	if err := collectRolls(ret, c); err != nil {
		return nil, err
	}

	if err := collectSeed(ret, c); err != nil {
		return nil, err
	}

	if err := collectKeep(ret, c); err != nil {
		return nil, err
	}

	if err := collectAggregation(ret, c); err != nil {
		return nil, err
	}

	if err := collectKeepCriteria(ret, c); err != nil {
		return nil, err
	}
	
	return ret, nil
}

func collectKeep(specification *roll.Specification, c *gin.Context) error {
	keep := c.Query("keep")
	if keep != "" {
		intKeep, err := strconv.Atoi(keep)
		if err != nil {
			return err
		}
		specification.Keep = intKeep
	}
	return nil
}

func collectSeed(specification *roll.Specification, c *gin.Context) error {
	seed := c.Query("seed")
	if seed != "" {
		intSeed, err := strconv.ParseInt(seed, 16, 64)
		if err != nil {
			return err
		}
		specification.Seed = intSeed
	}
	return nil
}

func collectRolls(specification *roll.Specification, c *gin.Context) error {
	rolls := c.Query("rolls")
	if rolls != "" {
		intRolls, err := strconv.Atoi(rolls)
		if err != nil {
			return err
		}
		specification.Rolls = intRolls
	}
	return nil
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
