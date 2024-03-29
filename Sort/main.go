package main

import (
	"fmt"

	"dsa/Sort/sort"
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
	fmt.Println("ORIGINAL ARRAY :-", arr)
	fmt.Println("BUBBLE SORT    :-", sort.BubbleSort(arr, limit))
	fmt.Println("INSERTION SORT :-", sort.InsertionSort(arr, limit))
	fmt.Println("SELECTION SORT :-", sort.SelectionSort(arr, limit))
	fmt.Println("QUICK SORT     :-", sort.QuickSort(arr, limit))
	fmt.Println("MERGE SORT     :-", sort.MergeSort(arr))

}
