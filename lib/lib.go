package lib

import (
	"strconv"
)

// Returns length of the longest string in the slice of strings.
func LongestStringLenInSlice(slice []string) int {
	maxLen := 0

	for _, s := range slice {
		sLen := len(s)
		if sLen > maxLen {
			maxLen = sLen
		}
	}

	return maxLen
}

func F64ToString(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}
