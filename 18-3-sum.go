package algorithms

// threeSum finds all unique triplets that sum up to 0 without using any imports.
func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	// 1. Sort the array in place using a simple Insertion Sort
	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	// 2. Iterate through the array, fixing the first element of the triplet
	for i := 0; i < n-2; i++ {
		// If the current element is greater than 0, the remaining elements
		// will also be greater than 0, so they can never sum to 0.
		if nums[i] > 0 {
			break
		}

		// Skip duplicate values for the first element to avoid duplicate triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Initialize two pointers
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// Found a valid triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate values for the second element
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicate values for the third element
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move both pointers inward after processing the current match
				left++
				right--
			} else if sum < 0 {
				// Sum is too small, move left pointer rightward to increase the sum
				left++
			} else {
				// Sum is too large, move right pointer leftward to decrease the sum
				right--
			}
		}
	}

	return result
}

// How It Works

// 1. Sorting (sort.Ints): Sorting the array takes $O(n \log n)$ time. It organizes the data so that duplicates sit next to each other, and it enables the two-pointer mechanic.
// 2. The Outer Loop (i): We pick nums[i] as our baseline first number. If nums[i] is ever greater than 0, we can instantly stop everything, because three positive numbers can never sum up to 0.
// 3. Avoiding Duplicates:For the outer loop, if i > 0 && nums[i] == nums[i-1] ensures we don't recalculate for the same starting number.Inside the pointer logic, nums[left] == nums[left+1] and nums[right] == nums[right-1] loops ensure we skip identical configurations after finding a successful triplet.
// 4. Two-Pointer Technique (left and right):If the sum is less than 0, we need a larger number, so we increment left.If the sum is greater than 0, we need a smaller number, so we decrement right.
