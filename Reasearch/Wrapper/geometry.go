package Wrapper

import (
	"fmt"
)

func main() {
	// Create a polygon with four vertices
	polygon := geometry.Polygon{
		{0, 0}, {0, 1}, {1, 1}, {1, 0},
	}

	// Check if a point is inside the polygon
	point := geometry.Point{0.5, 0.5}
	if polygon.Contains(point) {
		fmt.Println("Point is inside the polygon")
	} else {
		fmt.Println("Point is outside the polygon")
	}
}
