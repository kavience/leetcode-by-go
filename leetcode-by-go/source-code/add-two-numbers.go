package source_code

import "fmt"

// 1. 数组模式
func AddTwoNumbers(l1 []int, l2 []int) (sum []int) {
	var (
		minL    = l1
		maxL    = l2
		rest    []int
		tempSum = 0
		tempNum = 0
		flag    = 0
	)

	if len(l1) > len(l2) {
		minL = l2
		maxL = l1
	}
	rest = maxL[len(minL):]

	for i := 0; i < len(minL); i++ {
		tempSum = minL[i] + maxL[i] + flag
		tempNum = tempSum
		if tempSum >= 10 {
			tempNum = tempSum % 10
			flag = 1
		} else {
			flag = 0
		}
		sum = append(sum, tempNum)
	}

	if len(rest) == 0 {
		if flag > 0 {
			sum = append(sum, flag)
		}
	} else {
		for j := 0; j < len(rest); j++ {
			tempSum = rest[j] + flag
			tempNum = tempSum
			if tempSum >= 10 {
				tempNum = tempSum % 10
				flag = 1
			} else {
				flag = 0
			}
			sum = append(sum, tempNum)
		}

		if flag > 0 {
			sum = append(sum, flag)
		}
	}

	return sum
}

// 2. 链表方式
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoLinkNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headList := new(ListNode)
	head := headList
	num := 0

	for l1 != nil || l2 != nil || num > 0 {
		headList.Next = new(ListNode)
		headList = headList.Next
		if l1 != nil {
			num = num + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num = num + l2.Val
			l2 = l2.Next
		}
		headList.Val = (num) % 10
		num = num / 10
	}

	return head.Next
}

func LinkPrint(cur *ListNode) {
	lists := []int{}
	for cur != nil {
		lists = append(lists, cur.Val)
		cur = cur.Next
	}
	fmt.Println(lists)
}
