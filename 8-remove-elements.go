// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.

package algorithms

func removeElement(nums []int, val int) int {
	// k will track the number of elements not equal to val
	k := 0

	// Iterate through the slice with a read pointer 'i'
	for i := 0; i < len(nums); i++ {
		// If the current element is NOT the value we want to remove
		if nums[i] != val {
			// Move it to the 'k' position
			nums[k] = nums[i]
			// Increment k to prepare for the next valid element
			k++
		}
	}

	// k is the total count of elements that are not 'val'
	return k
}
