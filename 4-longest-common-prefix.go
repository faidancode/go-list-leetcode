// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"

package algorithms

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the initial prefix
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// Shorten the prefix until it matches the start of strs[i]
		for !strings.HasPrefix(strs[i], prefix) {
			if len(prefix) == 0 {
				return ""
			}
			// Slice the string to remove the last character
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

// To find the longest common prefix among an array of strings, the most efficient approach is "Horizontal Scanning." We assume the first string is the prefix and then compare it against each subsequent string, shortening it whenever a mismatch is found.

// Logic Breakdown

// Edge Case: If the input array is empty, return an empty string.

// Initialization: Start by assuming the first string (strs[0]) is the prefix.

// Iteration: Loop through the rest of the strings in the array.

// Comparison: For each string, check if it starts with the current prefix.
// - If it doesn't, remove the last character from the prefix and try again.
// - If the prefix becomes empty during this process, there is no common prefix, so return "".

// Return: After checking all strings, the remaining prefix is the longest common one.
