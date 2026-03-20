package algorithms

func IsPalindrome(x int) bool {
	// Special cases:
	// If x < 0, it is not a palindrome.
	// If the last digit of the number is 0, the first digit
	// of the number also needs to be 0 for it to be a palindrome.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// When the length is an odd number, we can get rid of the middle digit by revertedNumber/10
	// For example when x = 121, at the end of the loop we have x = 1, revertedNumber = 12
	// Since the middle digit doesn't matter in palindrome, we can simply get rid of it.
	return x == revertedNumber || x == revertedNumber/10
}

// Logic Breakdown
// Negative Numbers: Any negative number (e.g., -121) is not a palindrome because the negative sign becomes the last character when reversed.

// Trailing Zeros: If the last digit is 0, the first digit must also be 0 for it to be a palindrome (which only happens for the number 0 itself).

// Reversing Half: Instead of reversing the whole number (which might cause an integer overflow), we reverse the number digit by digit until the reversed part is greater than or equal to the remaining part.
