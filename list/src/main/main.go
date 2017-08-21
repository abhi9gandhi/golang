package main

import "fmt"
import "link_list"
import "os"

func main() {
	var option, data int = 0, 0
	var is_exit bool = false
	var list_type int = link_list.SINGLY
	for {
			fmt.Println("\nEnter the List type: 1- Singly 2- Doubly 3- Doubly Circular 4- Exit\n")
			fmt.Scanf("%d", &list_type)
		if list_type == 4 {
			os.Exit(0)
		}
		is_exit = false
		for {
			fmt.Println("\nEnter the option: 1- Add 2- Delete 3-Print 4- Exit\n")
			fmt.Scanf("%d", &option)

			switch option {
			case 1:
				fmt.Printf("Enter the data to add in list: ")
				fmt.Scanf("%d", &data)
				link_list.Add_callbacks[list_type](data)
				break
			case 2:
				fmt.Printf("Enter the data to be delete from list: ")
				fmt.Scanf("%d", &data)
				link_list.Delete_callbacks[list_type](data)
				break
			case 3:
				link_list.Print_callbacks[list_type]()
				break
			case 4:
				is_exit = true
				break
			}
			if is_exit == true {
				break;
			}
		}
	}
}
