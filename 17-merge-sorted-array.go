// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

package algorithms

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Index of the last element in the initial part of nums1
	i := m - 1
	// Index of the last element in nums2
	j := n - 1
	// Index of the last position in the entire nums1 array
	k := m + n - 1

	// While there are elements to compare in both arrays
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			// If the element in nums1 is larger, move it to the back
			nums1[k] = nums1[i]
			i--
		} else {
			// If the element in nums2 is larger or equal, move it to the back
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are remaining elements in nums2, copy them.
	// (Note: If there are remaining elements in nums1, they are already in place)
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
