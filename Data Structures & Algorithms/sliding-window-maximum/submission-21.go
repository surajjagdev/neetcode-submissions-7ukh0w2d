type Node struct {
	Prev *Node
	Next *Node
	Value int
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Listlen int
}

func NewLinkedList() *LinkedList{
	headNode := &Node{
		Prev: nil,
		Next: nil,
		Value: 0,
	}
	tailNode := &Node{
		Prev: nil,
		Next: nil,
		Value: 0,
	}

	headNode.Next = tailNode
	tailNode.Prev = headNode

	return &LinkedList{
		Head: headNode,
		Tail: tailNode,
		Listlen: 0,
	}
}

func (node *Node) breakLinks() {
	node.Next = nil
	node.Prev = nil
}

func (ll *LinkedList) Popfront() *Node{
	if ll.Head.Next == ll.Tail {
		return nil
	}

	// pop the front element
	node := ll.Head.Next

	node.Next.Prev = ll.Head
	ll.Head.Next = node.Next

	node.breakLinks()
	ll.Listlen--

	return node
}

func (ll *LinkedList) Append(v int) {
	// append to the back

	newNode := &Node{
		Prev: nil,
		Next: nil,
		Value: v,
	}
	lastNode := ll.Tail.Prev
	lastNode.Next = newNode
	ll.Tail.Prev = newNode

	newNode.Prev = lastNode
	newNode.Next = ll.Tail
	ll.Listlen++
}

func (ll *LinkedList) Popback() *Node {
	if ll.Head.Next == ll.Tail {
		return nil
	}

	lastNode := ll.Tail.Prev

	lastNode.Prev.Next = ll.Tail
	ll.Tail.Prev = lastNode.Prev

	lastNode.breakLinks()
	ll.Listlen--

	return lastNode
}

func (ll *LinkedList) Peekfront() int {
	if ll.Head.Next == ll.Tail {
		return -1
	}

	return ll.Head.Next.Value
}

func (ll *LinkedList) Peekback() int {
	if ll.Head.Next == ll.Tail {
		return -1
	}

	return ll.Tail.Prev.Value
}

func (ll *LinkedList) Len() int {
	return ll.Listlen
}



func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
	dq := NewLinkedList()
	left := 0

	for right := 0; right < len(nums); right++ {
		val := nums[right]

		// remove all elements from back of the queue if their value is less 
		// than new element's value
		lastIdx := dq.Peekback()

		// 1. maintain decreasing queue
		for lastIdx != -1 && nums[lastIdx] < val {
			rNode := dq.Popback()

			if rNode == nil {
				break
			}

			lastIdx = dq.Peekback()
		}

		// 2. add new element
		dq.Append(right)

		// check the front of our dq, is it out of range ?
		frontIdx := dq.Peekfront()

		// Remove out of window index
		if frontIdx != -1 && frontIdx < left {
			dq.Popfront()
		}

		// do we have enough elements for k ?
		if right >= k - 1 {
			res = append(res, nums[dq.Peekfront()])

			// increment left bc we keep moving 
			left++
		}

		
	}


	return res
}
