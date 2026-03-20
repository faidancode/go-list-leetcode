// File: algorithms/two_sum.go

package algorithms

func TwoSum(nums []int, target int) []int {
	// Create a map to store the value and its index
	// Key: the number, Value: the index of that number
	prevMap := make(map[int]int)

	for i, n := range nums {
		// Calculate the value needed to reach the target
		diff := target - n

		// Check if the complement already exists in our map
		if idx, found := prevMap[diff]; found {
			return []int{idx, i}
		}

		// Store the current number and its index in the map
		prevMap[n] = i
	}

	// Since the problem guarantees one solution, we don't
	// strictly need a return here, but Go requires it.
	return nil
}
