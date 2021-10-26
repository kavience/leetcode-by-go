package main

import (
	// "fmt"

	"fmt"
	source_code "leetcode-by-go/source-code"
	"time"
	// "leetcode-by-go/src"
)

func main() {
	var spendtime = time.Time{}
	/* 3sum.go */
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// spendtime =  time.Now()
	// fmt.Println(source_code.ThreeSum(nums))
	// fmt.Println(time.Since(spendtime))

	/* two-sum.go */
	// nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// spendtime = time.Now()
	// fmt.Println(source_code.MaxArea1(nums))
	// fmt.Println(time.Since(spendtime))

	// spendtime = time.Now()
	// fmt.Println(source_code.MaxArea2(nums))
	// fmt.Println(time.Since(spendtime))

	/* median-of-two-sorted-arrays.go */
	// nums1, nums2 := []int{1, 2, 3, 6, 9}, []int{2, 4, 5, 8}
	// spendtime = time.Now()
	// fmt.Println(source_code.FindMedianSortedArrays(nums1, nums2))
	// time.Since(spendtime)

	/* remove-duplicates-from-sorted-array.go */
	// nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// spendtime = time.Now()
	// fmt.Println(source_code.RemoveDuplicates(nums))
	// fmt.Println(time.Since(spendtime))

	/* remove-element.go */
	// nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// spendtime = time.Now()
	// fmt.Println(source_code.RemoveElement(nums))
	// fmt.Println(time.Since(spendtime))

	/* search-insert-position.go */
	// nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// spendtime = time.Now()
	// fmt.Println(source_code.SearchInsert(nums, 4))
	// fmt.Println(time.Since(spendtime))

	/* add-two-numbers.go */
	spendtime = time.Now()
	fmt.Println("array: ")
	l1 := []int{9, 9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9}
	result1 := source_code.AddTwoNumbers(l1, l2)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println("result1: ")
	fmt.Println(result1)
	fmt.Println(time.Since(spendtime))

	spendtime = time.Now()
	var (
		a = source_code.ListNode{Val: 9, Next: nil}
		b = source_code.ListNode{Val: 9, Next: &a}
		c = source_code.ListNode{Val: 9, Next: &b}
		d = source_code.ListNode{Val: 9, Next: &c}
		e = source_code.ListNode{Val: 9, Next: &d}
		f = source_code.ListNode{Val: 9, Next: &e}
	)
	l3 := f
	l4 := d
	result2 := source_code.AddTwoLinkNumbers(&l3, &l4)
	fmt.Println("list node: ")
	source_code.LinkPrint(&l3)
	source_code.LinkPrint(&l4)
	fmt.Println("result2: ")
	source_code.LinkPrint(result2)
	fmt.Println(time.Since(spendtime))
}
