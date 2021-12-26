package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	var head *node
	a := []int{1, 2, 3, 4, 5, 6}
	length := len(a)
	var s int
	fmt.Println("enter a number to add in the end")
	fmt.Scan(&s)
	
	b := make([]int,length+1)
	for i:= 0;i<length;i++{
		b[i] = a[i]
	}
	b[length] = s
	a = b
	// add(a,length)
	fmt.Println(a)
	for i:=0;i<length+1;i++ {
		n := node{data: a[i],}
		if head == nil {
			head = &n
			continue
		}
		last := head
		for last.next != nil {
			last = last.next
		}
		last.next = &n
	}
	for head != nil {
		
		fmt.Println(head.data,head.next)
		head = head.next
	}
}
// func add(a []int,length int){
	
// }