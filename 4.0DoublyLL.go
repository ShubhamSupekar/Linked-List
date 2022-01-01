package main

import "fmt"

type node struct {
	data   int
	before *node
	next   *node
}

func main() {
	var head *node
	var last *node
	var prev *node
	a := []int{10, 40, 60, 50, 20, 1}
	for i:=0;i<len(a);i++{
		n1 := node{data: a[i]}
		if head == nil {
			head = &n1
			last = head
			prev = head
			continue
		}
		temp:= &n1
		temp.before = prev
		prev = &n1
		last.next = &n1
		last = last.next
	}
	for head != nil {
		fmt.Println(head.before,head.data,head.next)
		head = head.next

	}
}