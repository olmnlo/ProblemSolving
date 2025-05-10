package vjodge

import "fmt"

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
		totalSum += digits[i]
	}

	found := false

	for minimumSegments := 2; minimumSegments <= n; minimumSegments++ {
		if totalSum%minimumSegments != 0 {
			continue
		}

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
