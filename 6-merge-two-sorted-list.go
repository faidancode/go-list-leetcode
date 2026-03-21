package algorithms

// ListNode mendefinisikan struktur untuk linked list.
// type ListNode struct {
//     Val  int
//     Next *ListNode
// }

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Membuat dummy node sebagai titik awal
	dummy := &ListNode{}
	current := dummy

	// Iterasi selama kedua list masih memiliki node
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		// Gerakkan pointer current ke depan
		current = current.Next
	}

	// Jika salah satu list sudah habis, sambungkan sisanya
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	// Return mulai dari node setelah dummy
	return dummy.Next
}
