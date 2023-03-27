package Reasearch

import (
"fmt"
"github.com/twpayne/go-geom/encoding/geojson"
)

func main() {
	// Create a GeoJSON polygon with four vertices
	coords := [][]float64{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
		{0, 0},
	}
	polygon := geojson.NewPolygonGeometry([][][]float64{coords})

	// Check if a point is inside the polygon
	point := geojson.NewPointGeometry([]float64{0.5, 0.5}

}

