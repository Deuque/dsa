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
		sum = intFromString(splitStringAndReverse(*d2))
	} else if d2 == nil {
		sum = intFromString(splitStringAndReverse(*d1))
	} else {
		sum = intFromString(splitStringAndReverse(*d1)) + intFromString(splitStringAndReverse(*d2))
	}

	fmt.Println(sum)
	sumsplit := strings.Split(fmt.Sprintf("%v", sum), "")
	rsumsplit := reverse(sumsplit)
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
func splitStringAndReverse(item string) string {
	fmt.Println(item)
	split := strings.Split(item, "")
	rsplit := reverse(split)
	return strings.Join(rsplit, "")

}

func reverse(numbers []string) []string {
	newNumbers := make([]string, len(numbers))
	for i, j := 0, len(numbers)-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
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
