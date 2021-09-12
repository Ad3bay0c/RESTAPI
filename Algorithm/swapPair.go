//https://leetcode.com/problems/swap-nodes-in-pairs/submissions/

package Algorithm

type ListNode struct {
    Val int
    Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	current := head
	first, second := &ListNode{}, &ListNode{}

	count := 1

	for current != nil {
		if count % 2 == 0 {
			if count == 2 {
				head = first.Next
				first.Next = current.Next
				head.Next = first
				current = head.Next
				// second.Next = head.Next
			} else {
				first.Next = current.Next
				current.Next = first
				second.Next = current
				current = current.Next
			}
			second = current
		} else {
			first = current
		}

		count++
		current = current.Next
	}

	return head

}
