package utils

import "math"

// TruncateToSixDecimalPlaces : truncate float64 to six decimal places
func TruncateToSixDecimalPlaces(f float64) float64 {
	return math.Round(f*1e6) / 1e6
}
