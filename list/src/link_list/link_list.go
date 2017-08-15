package link_list

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

var head *node = nil

func Add_list(data int) {
	var iter *node = head
	
	new_node := new(node)
	new_node.data = data
	new_node.next = nil
		
	if head == nil {
		head = new_node
	} else {	
		for iter.next != nil {
			iter = iter.next
		}
		iter.next = new_node
	}
   return 		
}

func Delete_list(data int) {
	var iter *node = head
	var is_found bool = false
	
	if head == nil {
		fmt.Printf("List is empty")
		return
	}
	
	if head.data == data {
		head = head.next
		return
	}
	
	for iter.next != nil {
		if iter.next.data == data {
			is_found = true
			iter.next = iter.next.next
			break
		}
		iter = iter.next
	}
	
	if is_found == false {
		fmt.Printf("Could not find element %d in existing list", data)
	}
	
	return
}

func Print_list() {
	var iter *node = head
	
	fmt.Printf("\nList data: \n")
	if head == nil {
		return
	} 
	
	for iter != nil {
		fmt.Printf(" %d ", iter.data)
		iter = iter.next
	}
	return
} 

