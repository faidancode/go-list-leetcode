package algorithms

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	// If needle is longer than haystack, it can't be a match
	if m > n {
		return -1
	}

	// Iterate through the haystack up to the point where needle can still fit
	for i := 0; i <= n-m; i++ {
		// Check if the substring starting at i matches needle
		// - i: The starting index (inclusive).
		// - i+m: The ending index (exclusive).
		// - The Length: The resulting substring will always have a length of exactly m (because (i + m) - i = m).
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}
