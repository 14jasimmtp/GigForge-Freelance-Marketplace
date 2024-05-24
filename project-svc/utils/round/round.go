package round

import "math"

func RoundToTwoDecimalPlaces(number float64) float64 {
	factor := math.Pow(10, 2)
	return math.Round(number*factor) / factor
}
