// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

package algorithms

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		// Calculate mid-point (using low + (high-low)/2 to prevent overflow)
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// If not found, 'low' is the insertion index
	return low
}

// How it Works
// Divide and Conquer: We check the middle element.
// - If nums[mid] is the target, we are done.
// - If the target is larger, we ignore the left half (low = mid + 1).
// - If the target is smaller, we ignore the right half (high = mid - 1).

// The "Invisible" Insertion Point: When the loop ends (low > high), it means the target wasn't in the array. At this specific moment, the low pointer has moved past all elements smaller than the target and is pointing exactly where the target belongs.

// Example Walkthrough
// Input: nums = [1, 3, 5, 6], target = 2
// Initial: low = 0, high = 3. mid = 1 (nums[1] = 3).
// Compare: 3 > 2, so high = mid - 1 (which is 0).
// Next Loop: low = 0, high = 0. mid = 0 (nums[0] = 1).
// Compare: 1 < 2, so low = mid + 1 (which is 1).
// Exit: low (1) is no longer <= high (0).
// Result: Return low, which is 1. (Correct: [1, 2, 3, 5, 6]).
