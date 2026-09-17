package lib

// Returns length of the longest string in the slice of strings.
func LongestStringInSlice(slice []string) int {
	maxLen := 0

	for _, s := range slice {
		sLen := len(s)
		if sLen > maxLen {
			maxLen = sLen
		}
	}

	return maxLen
}
