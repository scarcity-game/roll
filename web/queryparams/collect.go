package queryparams

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const minParam = "min"
const maxParam = "max"
const meanParam = "mean"
const stddevParam = "stddev"

func collectFloat64(param string, c *gin.Context, defaultValue float64) (float64, error) {
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

func collectInt(param string, c *gin.Context, defaultValue int) (int, error) {
	found := c.Query(param)
	if found != "" {
		asInt, err := strconv.Atoi(found)
		if err != nil {
			return 0, err
		}
		return asInt, nil
	}
	return defaultValue, nil
}

func collectInt64(param string, c *gin.Context, defaultValue int64) (int64, error) {
	found := c.Query(param)
	if found != "" {
		asInt, err := strconv.ParseInt(found, 16, 64)
		if err != nil {
			return 0, err
		}
		return asInt, nil
	}
	return defaultValue, nil
}
