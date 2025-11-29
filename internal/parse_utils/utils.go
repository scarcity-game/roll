package parse_utils

import (
	"strconv"
	"time"
)

func ParseSeed(seedString string) (int64, error) {
	if seedString != "" {
		asInt, err := strconv.ParseInt(seedString, 16, 64)
		if err != nil {
			return 0, err
		}
		return asInt, nil
	}
	return time.Now().UnixNano(), nil
}
