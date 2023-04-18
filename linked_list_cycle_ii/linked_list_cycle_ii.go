package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution to 142

func detectCycle(head *ListNode) *ListNode {
	set := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := set[head]; ok {
			return head
		}
		set[head] = struct{}{}
		head = head.Next
	}

	return nil

}
