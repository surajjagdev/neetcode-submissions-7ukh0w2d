/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // need half way point
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// slow is at mid point

	// reverse second half
	// cap first half
	curr := slow.Next
	slow.Next = nil
	var prev *ListNode = nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	p1 := head
	p2 := prev

	for p2 != nil {
		tmp1 := p1.Next
		tmp2 := p2.Next

		p1.Next = p2
		p2.Next = tmp1

		p1 = tmp1
		p2 = tmp2
	}
}
