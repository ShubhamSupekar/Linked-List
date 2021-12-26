package main

import "fmt"

type node struct {
	data int
	next *node
}
var head *node = new(node)
func info() {
	
	head.data = 1
	head.next = nil
	var value1 *node = &node{data: 2, next: nil}
	head.next = value1

	var value2 *node = &node{data: 3, next: nil}
	value1.next = value2

	var tail *node = &node{data: 4, next: nil}
	value2.next = tail
}

func main(){
	info()
	for {
		if head == nil {
			break
		}
		fmt.Printf("\n%d,%d\n",head.data,head.next)
		fmt.Printf("%d ->",head.data)
		head = head.next
	}
	fmt.Printf("\nEND\n")

}