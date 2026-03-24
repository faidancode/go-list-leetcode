package algorithms

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			// Successful match found, return early!
			return digits
		}
	}

	// This part only runs for cases like [9, 9, 9]
	result := make([]int, n+1)
	result[0] = 1
	return result
}
