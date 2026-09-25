package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(distance(55.7558, 37.6173, 59.9343, 30.3351))
}

func distance(lat1, lon1 float32, lat2, lon2 float32) float32 {
	lat1r, lat2r, lon1r, lon2r := getRadian(lat1), getRadian(lat2), getRadian(lon1), getRadian(lon2)
	deltaLat := lat2r - lat1r
	deltaLon := lon2r - lon1r

	a := math.Pow(math.Sin(deltaLat/2), 2)
	b := math.Cos(lat1r)
	c := math.Cos(lat2r)
	d := math.Pow(math.Sin(deltaLon/2), 2)
	x := b * c * d
	sqr := math.Sqrt(a + x)
	res := 2 * 6371 * math.Asin(sqr)
	return float32(res)
}

func getRadian(point float32) float64 {
	return float64(point) * (math.Pi / 180)
}
