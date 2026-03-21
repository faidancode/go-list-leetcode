package algorithms

func RomanToInt(s string) int {
	// Map symbols to their integer values
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currentVal := romanMap[s[i]]

		// Check if there is a next symbol and if it's larger than the current one
		if i+1 < n && currentVal < romanMap[s[i+1]] {
			// Subtraction rule: subtract current value
			total -= currentVal
		} else {
			// Addition rule: add current value
			total += currentVal
		}
	}

	return total
}
