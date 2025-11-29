package parse_utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func CollectFloat64(param string, c *gin.Context, defaultValue float64) (float64, error) {
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

func CollectInt(param string, c *gin.Context, defaultValue int) (int, error) {
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
