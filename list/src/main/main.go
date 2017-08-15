package main
import "fmt"
import "link_list"

func main() {
	var option, data int = 0, 0
	for {
		fmt.Println("\nEnter the option: 1- Add 2- Delete 3-Print\n")
		fmt.Scanf("%d", &option)
		
		switch option {
			case 1:
				fmt.Printf("Enter the data to add in list: ")
				fmt.Scanf("%d", &data)
				link_list.Add_list(data)
			break
			case 2:
				fmt.Printf("Enter the data to be delete from list: ")
				fmt.Scanf("%d", &data)
				link_list.Delete_list(data)
			break
			case 3:
				link_list.Print_list()
			break
		}		
	}	
}