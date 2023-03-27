package Reasearch

import (
	"fmt"
)

func main() {
	// Create a quad tree index and insert points
	index := rquadtree.NewIndex(0, 0, 1, 1)
	index.Insert(rquadtree.NewPoint(0.25, 0.25))
	index.Insert(rquadtree.NewPoint(0.75, 0.75))

	// Check if a point is inside the bounding box of any inserted point
	point := rquadtree.NewPoint(0.5, 0.5)
	found := index.Search(point, func(item rquadtree.Item) bool {
		return item.BBox().Contains(point)
	})
	if len(found) > 0 {
		fmt.Println("Point is inside a bounding box")
	} else {
		fmt.Println("Point is outside all bounding boxes")
	}
}
