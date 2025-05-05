package leetcode

import (
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	charIndex := make(map[string]int)
	maxLen := 0
	start := 0
	my := strings.Split(s, "")

	for i, ch := range my {
		if lastIndex, found := charIndex[ch]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndex[ch] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}

	return maxLen
}
