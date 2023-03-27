package Reasearch

import (
	"fmt"
)

func main() {
	// Create a polygon with four vertices
	coords := [][]float64{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
		{0, 0},
	}
	polygon := geojson.NewPolygonGeometry([][][]float64{coords})

	// Check if a point is inside the polygon
	point := geojson.NewPointGeometry([]float64{0.5, 0.5})
	if polygon.Contains(point) {
		fmt.Println("Point is inside the polygon")
	} else {
		fmt.Println("Point is outside the polygon")
	}
}
