package vjodge

import (
	"fmt"
	"log"
)

func VasyaGoldenTicket() {
	var n int
	var digitsStr string
	// fmt.Scan(&n)
	// fmt.Scan(&digitsStr)
	n = 5
	digitsStr = "12345"

	digits := make([]int, n)
	totalSum := 0

	for i, c := range digitsStr {
		digits[i] = int(c - '0')
		log.Printf("loop: %d, the array: %v\n", i, digits)
		totalSum += digits[i]
		log.Printf("the total sum in: %d\n", totalSum)
	}

	found := false

	for minimumSegments := 2; minimumSegments <= n; minimumSegments++ {
		if totalSum%minimumSegments != 0 {
			log.Printf("inside (IF) the loop K: %d, totalsum: %d", minimumSegments, totalSum)
			continue
		}
		log.Printf("the loop K: %d, totalsum: %d", minimumSegments, totalSum)

		s := totalSum / minimumSegments
		currentSum := 0
		segments := 0
		for _, num := range digits {
			currentSum += num
			if currentSum == s {
				segments++
				currentSum = 0
			} else if currentSum > s {
				break
			}
			if segments == minimumSegments && currentSum == 0 {
				found = true
				break
			}
		}
	}

	if found {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
