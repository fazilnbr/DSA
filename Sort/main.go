package main

import (
	"fmt"

	"sort.go/sort"
)

func main() {
	var limit int
	fmt.Print("\nEnter the limit of array :-")
	fmt.Scan(&limit)
	arr := make([]int, limit)
	fmt.Println("Enter the values:- ")
	for i := 0; i < limit; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("BUBBLE SORT    :-", sort.BubbleSort(arr, limit))
	fmt.Println("INSERTION SORT :-", sort.InsertionSort(arr, limit))
	fmt.Println("SELECTION SORT :-", sort.SelectonSort(arr, limit))
}
