package algorithms

func mySqrt(x int) int {
	// Base cases for 0 and 1
	if x < 2 {
		return x
	}

	left := 1
	right := x / 2
	ans := 0

	for left <= right {
		// Calculate middle point to avoid potential overflow: mid = left + (right - left) / 2
		mid := left + (right-left)/2

		// Use division to check: mid <= x/mid (equivalent to mid*mid <= x)
		// This prevents overflow if x was near the maximum limit of int32/int64
		if mid <= x/mid {
			// mid is a potential candidate, so we store it
			ans = mid
			// Look for a larger potential square root in the right half
			left = mid + 1
		} else {
			// mid is too large, look in the left half
			right = mid - 1
		}
	}

	return ans
}

// Why Binary Search for Square Roots?
// The problem asks for an integer $k$ such that $k^2 \le x < (k+1)^2$. Because the function $f(k) = k^2$ is monotonically increasing for non-negative integers, the numbers are already "sorted" in terms of their squares. This makes binary search the most efficient $O(\log n)$ approach compared to a linear $O(n)$ scan.

// Key Details in the Go code:
// - Integer Division: Since we are returning an integer and rounding down, Go’s default integer division / handles the "floor" requirement naturally.
// - Overflow Prevention: Even though Go's int is usually 64-bit, using mid <= x/mid instead of mid*mid <= x is a best-practice "defensive programming" habit to prevent exceeding the maximum value of a variable type.
// - Range Optimization: Starting right at x/2 is a safe optimization because $\sqrt{x} \le x/2$ for all $x \ge 4$.
