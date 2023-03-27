package Reasearch

import (
	"fmt"
	"github.com/golang/geo/s2"
)

func main() {
	// Create a loop (polygon) with four vertices
	loop := s2.LoopFromPoints([]s2.Point{
		s2.PointFromLatLng(s2.LatLngFromDegrees(0, 0)),
		s2.PointFromLatLng(s2.LatLngFromDegrees(0, 1)),
		s2.PointFromLatLng(s2.LatLngFromDegrees(1, 1)),
		s2.PointFromLatLng(s2.LatLngFromDegrees(1, 0)),
	})

	// Check if a point is inside the loop
	point := s2.PointFromLatLng(s2.LatLngFromDegrees(0.5, 0.5))
	if loop.ContainsPoint(point) {
		fmt.Println("Point is inside the loop")
	} else {
		fmt.Println("Point is outside the loop")
	}
}
