package pointers

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func NewLinkedListNode(data int, next *LinkedListNode) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = nil
	return node
}

func removeNthLastNode(head *LinkedListNode, n int) *LinkedListNode {

	//Put pointer on the head of the linked list
	right := head
	left := head

	// move right pointer by nth elements
	for i := 0; i < n; i++ {
		right = right.next
	}

	if right == nil {
		return head.next
	}
	//move pointers one step
	for right.next != nil {
		right = right.next
		left = left.next
	}
	// it is right above the node to be deleted
	left.next = left.next.next

	return head
}
