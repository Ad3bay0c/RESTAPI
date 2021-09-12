//https://leetcode.com/problems/merge-k-sorted-lists/

package Algorithm

import "sort"

func mergeKLists(lists []*ListNode) *ListNode {
	var arrayLists = make([]int, 0)

	for _, val := range lists {
		current := val
		for current != nil {
			arrayLists = append(arrayLists, current.Val)
			current = current.Next
		}
	}

	sort.Ints(arrayLists)

	NewLists := &ListNode{}
	dummy := NewLists
	for _, val := range arrayLists {
		dummy.Next = &ListNode{Val: val}
		dummy = dummy.Next
	}

	return NewLists.Next
}
