package atc

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d1 := getAllDigits(l1)
	d2 := getAllDigits(l2)
	var sum int
	if d1 == nil && d2 == nil {
		return nil
	}
	if d1 == nil {
		sum = intFromString(reverseString(*d2))
	} else if d2 == nil {
		sum = intFromString(reverseString(*d1))
	} else {
		sum = intFromString(reverseString(*d1)) + intFromString(reverseString(*d2))
	}

	fmt.Println(sum)
	sumsplit := strings.Split(fmt.Sprintf("%v", sum), "")
	rsumsplit := reverseList(sumsplit)
	fmt.Println(rsumsplit)
	return appendAllDigits(rsumsplit)

}

func intFromString(item string) int {
	val, _ := strconv.Atoi(item)
	return val
}
func FloatFromString(item string) float64 {
	val, _ := strconv.ParseFloat(item, 64)
	return val
}

func reverseList(numbers []string) []string {
	return strings.Split(reverseString(strings.Join(numbers, "")), "")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getAllDigits(l *ListNode) *string {
	if l == nil {
		return nil
	}
	comb := fmt.Sprintf("%v", l.Val)
	if l.Next != nil {
		result := getAllDigits(l.Next)
		if result != nil {
			comb = fmt.Sprintf("%v%v", comb, *result)
		}

	}
	return &comb
}

func appendAllDigits(items []string) *ListNode {
	if len(items) == 0 {
		return nil
	}
	val, err := strconv.Atoi(items[0])
	if err != nil {
		return nil
	}

	others := appendAllDigits(items[1:])

	//fmt.Printf("%v : %v\n", items, others)
	return &ListNode{
		Val:  val,
		Next: others,
	}
}
