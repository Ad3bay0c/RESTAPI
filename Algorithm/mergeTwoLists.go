//https://leetcode.com/problems/merge-two-sorted-lists/

package Algorithm

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	current := l1
	current2 := l2

	var Node = &ListNode{}
	dummy := Node
	for {
		if current == nil {
			dummy.Next = current2
			break
		}
		if current2 == nil {
			dummy.Next = current
			break
		}

		if current.Val > current2.Val {

			dummy.Next = current2

			current2 = current2.Next

		} else {
			dummy.Next = current
			current = current.Next

		}

		dummy = dummy.Next
	}

	return Node.Next

	//     var lists = make([]int, 0)

	//     current := l1
	//     for current != nil {
	//         lists = append(lists, current.Val)
	//         current = current.Next
	//     }

	//     current = l2
	//     for current != nil {
	//         lists = append(lists, current.Val)
	//         current = current.Next
	//     }

	//     sort.Ints(lists)

	//     var newLinkedLists, currentVal *ListNode

	//     for i, val := range lists {
	//         if i  == 0 {
	//             newLinkedLists =  &ListNode{Val: val}
	//             currentVal = newLinkedLists
	//         } else {
	//             currentVal.Next = &ListNode{Val: val}
	//             currentVal = currentVal.Next
	//         }
	//     }
	//     return newLinkedLists
}