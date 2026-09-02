package lib

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
