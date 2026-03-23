// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// - Open brackets must be closed by the same type of brackets.
// - Open brackets must be closed in the correct order.
// - Every close bracket has a corresponding open bracket of the same type

package algorithms

func isValid(s string) bool {
	// A slice to act as our stack
	stack := []rune{}

	// Map for matching pairs: key is the closing bracket, value is the opening bracket
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		// If the character is a closing bracket

		// The "Comma OK" Idiom
		// In Go, when you access a map, it can return two values.
		// -opener: The value associated with the key (e.g., if char is ']', opener becomes '[').
		// -isClosing: A boolean that is true if the key exists in the map, and false if it doesn't.

		if opener, isClosing := pairs[char]; isClosing {
			// Check if stack is empty or the top doesn't match the opener
			if len(stack) == 0 || stack[len(stack)-1] != opener {
				return false
			}
			// Pop the element from the stack
			// In Go, we "pop" from a slice by re-slicing it.
			// We tell the stack: "Your new content is everything from the beginning up to (but not including) the last element."
			stack = stack[:len(stack)-1]
		} else {
			// It's an opening bracket, push it onto the stack
			stack = append(stack, char)
		}
	}

	// Return true only if the stack is empty
	return len(stack) == 0
}
