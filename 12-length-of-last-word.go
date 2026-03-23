// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal substring consisting of non-space characters only.

// Example 1:

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.

package algorithms

func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1

	// Step 1: Skip all spaces at the end of the string
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// Step 2: Count characters until we hit a space or the start of the string
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

// Step-by-Step Walkthrough
// Input: s = "   fly me   to   the moon  "
// Initial State: i = 26 (pointing at the last space).
// Skip Spaces: The first for loop moves i back.
// s[26] is ' ' $\rightarrow$ i = 25
// s[25] is ' ' $\rightarrow$ i = 24
// s[24] is 'n'. Stop.
// Count Word: The second for loop starts at i = 24.
// s[24] ('n'): length = 1, i = 23
// s[23] ('o'): length = 2, i = 22
// s[22] ('o'): length = 3, i = 21
// s[21] ('m'): length = 4, i = 20
// Stop: s[20] is a space ' '. The loop ends.
// Result: Return 4.
