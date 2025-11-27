package gaussian

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ExtractGaussianSpecification(c *gin.Context) (*Specification, error) {
	ret := defaultSpecification()
	if err := collectMean(ret, c); err != nil {
		return nil, err
	}

	if err := collectStddev(ret, c); err != nil {
		return nil, err
	}

	if err := collectMin(ret, c); err != nil {
		return nil, err
	}
	if err := collectMax(ret, c); err != nil {
		return nil, err
	}
	return ret, nil
}

func collectMean(specification *Specification, c *gin.Context) error {
	mean, err := collect("mean", c, specification.mean)
	if err != nil {
		return err
	}
	specification.mean = mean
	return nil
}
func collectStddev(specification *Specification, c *gin.Context) error {
	stddev, err := collect("stddev", c, specification.stddev)
	if err != nil {
		return err
	}
	specification.stddev = stddev
	return nil
}
func collectMin(specification *Specification, c *gin.Context) error {
	m, err := collect("min", c, specification.min)
	if err != nil {
		return err
	}
	specification.min = m
	return nil
}
func collectMax(specification *Specification, c *gin.Context) error {
	m, err := collect("max", c, specification.max)
	if err != nil {
		return err
	}
	specification.max = m
	return nil
}

func collect(param string, c *gin.Context, defaultValue float64) (float64, error) {
	found := c.Query(param)
	if found != "" {
		asFloat, err := strconv.ParseFloat(found, 64)
		if err != nil {
			return 0, err
		}
		return asFloat, nil
	}
	return defaultValue, nil
}
