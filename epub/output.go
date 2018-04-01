package epub

import (
	"fmt"
	"log"
)

// Output generates one file from all epub archieve
func (b *Book) Output() {

	for _, v := range b.Opf.Spine.Items {
		log.Println(v)
	}
	for _, v := range b.Opf.Manifest {
		fmt.Println(v)
	}
}

func deepThroughContent(points []NavPoint) []NavPoint {
	result := make([]NavPoint, 0)
	for _, v := range points {
		result = append(result, v.Points...)
		result = append(result, deepThroughContent(v.Points)...)
	}

	return result
}
