package link_list

import "fmt"

const (
	SINGLY = 1
	DOUBLY = 2
	DOUBLY_CIRCULAR = 3
)

type node struct {
	data int
	next *node
	prev *node
}

var head *node = nil


type add_list func (data int)
type delete_list func (data int)
type print_list func ()

var Add_callbacks = map[int] add_list {1: add_singly_list, 2: add_doubly_list, 3: add_doubly_circular_list }
var Delete_callbacks = map[int] delete_list {1: delete_singly_list, 2: delete_doubly_list, 3: delete_doubly_circular_list }
var Print_callbacks = map[int] print_list {1: print_singly_list, 2: print_doubly_list, 3: print_doubly_circular_list }

func add_singly_list(data int) {
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

func add_doubly_list(data int) {
	var iter *node = head
	
	new_node := new(node)
	new_node.data = data
	new_node.next = nil
	new_node.prev = nil
		
	if head == nil {
		head = new_node
	} else {	
		for iter.next != nil {
			iter = iter.next
		}
		iter.next = new_node
		new_node.prev = iter
	}
   return 		
}

func add_doubly_circular_list(data int) {
	var iter *node = head
	
	new_node := new(node)
	new_node.data = data
		
	if head == nil {
		head = new_node
		new_node.next = head
		new_node.prev = head
	} else {	
		for iter.next != head {
			iter = iter.next
		}
		
		new_node.next = iter.next
		iter.next.prev = new_node
		iter.next = new_node
		new_node.prev = iter
	}
   return 		
}

func delete_singly_list(data int) {
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

func delete_doubly_list(data int) {
	var iter *node = head
	var is_found bool = false
	
	if head == nil {
		fmt.Printf("List is empty")
		return
	}
	
	if head.data == data {
		if head.next != nil {
			head.next.prev = nil
		}
		head = head.next
		return
	}
	
	iter = iter.next
	
	for iter != nil {
		if iter.data == data {
			is_found = true
			if iter.next != nil {
				iter.next.prev = iter.prev
			} 
			
			iter.prev.next = iter.next
			break
		}
	
		iter = iter.next
	}
	
	if is_found == false {
		fmt.Printf("Could not find element %d in existing list", data)
	}
	
	return
}

func delete_doubly_circular_list(data int) {
	var iter *node = head
	var is_found bool = false
	
	if head == nil {
		fmt.Printf("List is empty")
		return
	}
	
	if head.data == data {
		if head.next == head {
			head = nil
			return
		}
		head.prev.next = head.next
		head.next.prev = head.prev
		head = head.next
		return
	}
	
	iter = head.next
	
	for iter != head {
		if iter.data == data {
			is_found = true
			
			iter.next.prev = iter.prev
			iter.prev.next = iter.next
			break
		}
		iter = iter.next
	}
	
	if is_found == false {
		fmt.Printf("Could not find element %d in existing list", data)
	}
	
	return
}

func print_singly_list() {
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

func print_doubly_list() {
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

func print_doubly_circular_list() {
	var iter *node = head
	
	fmt.Printf("\nList data: \n")
	if head == nil {
		return
	} 
	
	for iter.next != head {
		fmt.Printf(" %d ", iter.data)
		iter = iter.next
	}
	
	fmt.Printf(" %d ", iter.data)
	
	return
}