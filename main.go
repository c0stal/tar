package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func printList(s string, l *ListNode) {
	fmt.Print(s, ": ")
	for {
		if l == nil {
			fmt.Println("")
			return
		}
		fmt.Print(l.Val, ",")
		l = l.Next
	}
}
