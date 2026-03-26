// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

// Input: head = [1,1,2]
// Output: [1,2]

package algorithms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * Val int
 * Next *ListNode
 * }
 */

// DeleteDuplicates removes duplicate values from a sorted linked list.
// Time Complexity: O(n) - We visit each node exactly once.
// Space Complexity: O(1) - we transform the list in-place.
func DeleteDuplicates(head *ListNode) *ListNode {
	// 1. Guard Clause: If the list is empty or has only one node, no duplicates possible.
	if head == nil || head.Next == nil {
		return head
	}

	// 2. Use a pointer to traverse the list.
	// 'current' starts at the beginning.
	current := head

	// 3. Loop until we reach the end of the list.
	for current != nil && current.Next != nil {
		// Compare current node value with the next node value.
		if current.Val == current.Next.Val {
			// Found a duplicate!
			// Skip the next node by pointing 'Next' to the node after it.
			current.Next = current.Next.Next
		} else {
			// No duplicate, just move the pointer forward.
			current = current.Next
		}
	}

	// 4. Return the modified head of the list.
	return head
}

//Best Practices & Readability Analysis

// 1. Performance: In-Place Manipulation
// This solution is $O(n)$ because we touch every node once. More importantly, it is $O(1)$ space. We don't create a new list or use a hash map to track seen values because the Sorted constraint allows us to only look at the immediate neighbor.

// 2. Readability: Pointer Logic
// We use current := head to maintain a reference to the start of the list. Since head itself never changes (the first element of a sorted list will always be the "smallest" and thus never deleted as a duplicate of something before it), we can safely return head at the end.

// 3. Defensive Programming (Edge Cases)
// The check if head == nil || head.Next == nil is crucial. Without it, accessing current.Next.Val inside the loop could trigger a segmentation fault (panic) in Go if the list is empty.

// 4. The "Skip" LogicNotice that when we delete a duplicate:

// Gocurrent.Next = current.Next.Next

// We do not move the current pointer forward immediately. We stay on the same node to check if the new current.Next is also a duplicate (e.g., in a list like [1, 1, 1]). We only move forward in the else block when we are sure the next node is different.
