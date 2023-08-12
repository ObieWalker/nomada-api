package utils

import (
  "strconv"
	"time"
)

func ConvertToInt(stringValue string) int {
	value, _ := strconv.Atoi(stringValue)
	return value
}

func ConvertToTime(stringValue string) time.Duration {
	value, _ := time.ParseDuration(stringValue)
	return value
}

