package vjodge

import (
	"fmt"
)

func Battery() {
	var t int
	fmt.Scan(&t)
	levels := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Scan(&levels[i])
	}

	for _, battery := range levels {
		count := 0
		visited := make(map[int]bool) // Prevent infinite loops

		for battery != 50 {
			if visited[battery] {
				// No way to reach exactly 50%
				count = -1
				break
			}
			visited[battery] = true

			if battery < 50 {
				battery += 2
			} else {
				battery -= 3
			}

			if battery < 0 || battery > 100 {
				count = -1
				break
			}
			count++
		}

		if count == -1 {
			fmt.Println("Impossible")
		} else {
			fmt.Println(count)
		}
	}
}
