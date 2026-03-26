// You are climbing a staircase. It takes n steps to reach the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

// Example 1:
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps

package algorithms

// ClimbStairs calculates the number of distinct ways to reach the top.
// Time Complexity: O(n) - We loop through the steps once.
// Space Complexity: O(1) - We only store two variables regardless of n.
func ClimbStairs(n int) int {
	// 1. Handle edge cases immediately for readability.
	// If n is 1 or 2, the answer is n itself.
	if n <= 2 {
		return n
	}

	// 2. Use descriptive variable names that reflect the logic.
	// 'prev' is ways to (i-2), 'current' is ways to (i-1).
	prev := 1    // Ways to reach step 1
	current := 2 // Ways to reach step 2

	// 3. Start loop from the first "unknown" step.
	for i := 3; i <= n; i++ {
		// The ways to reach the current step is the sum of the previous two steps.
		next := prev + current

		// 4. Advance the window.
		prev = current
		current = next
	}

	return current
}
