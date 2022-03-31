package main

import "fmt"

func main() {
	// Input 10 integer
	var list = make([]int, 10)
	for i := 0; i < len(list); i++ {
		fmt.Print("Enter a number: ")
		fmt.Scan(&list[i])
	}

	// Print list with BubbleSort
	fmt.Println("List before BubbleSort: ", list)
	fmt.Println("List after BubbleSort: ", BubbleSort(list))
}

// BubbleSort
func BubbleSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				swap(list, j)
			}
		}
	}
	return list
}

// swap function
func swap(list []int, pos int) {
	list[pos], list[pos+1] = list[pos+1], list[pos]
}
