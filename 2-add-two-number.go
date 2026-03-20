// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

package algorithms

// ListNode defines a node for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// dummyHead allows us to easily return the start of the result list
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	carry := 0

	// Continue as long as there are nodes to process or a remaining carry
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Calculate new carry and the digit to store
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}

		// Move the pointer forward
		current = current.Next
	}

	return dummyHead.Next
}

// Key Logic
// BreakdownThe Dummy Head: Use dummyHead to simplify the linked list construction. It acts as a placeholder so we don't have to write extra logic to handle the head of the list separately.
// The Loop Condition: The loop runs as long as l1 or l2 has nodes, OR if there is a leftover carry (like in Example 3 where the final addition creates a new digit).
// Basic Math:
// - The value for the new node is $sum \pmod{10}$.
// - The carry for the next iteration is $\lfloor sum / 10 \rfloor$.
// Memory Management: Each iteration creates a new ListNode containing the result of that specific decimal position.
