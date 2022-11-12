package main

import (
	"fmt"
	"search/search"
	"time"
)

func main() {
	var size, target int
	fmt.Println("enter the size :-")
	fmt.Scan(&size)
	arr := make([]int, size)
	fmt.Println("enter the values : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("enter the number to search :-")
	fmt.Scan(&target)
	t := time.Now()
	res, pos := search.LeaniarSearch(target, arr)
	if res {
		fmt.Println("leaniar search\nthe velue found on the posisson : ", pos)
	} else {
		fmt.Println("Not founded!!!!")
	}
	fmt.Println(time.Since(t))

	t1 := time.Now()
	res1, pos1 := search.BainarySearch(target, arr)
	if res1 {
		fmt.Println("binary search\nthe velue found on the posisson : ", pos1)
	} else {
		fmt.Println("Not founded!!!!")
	}
	fmt.Println(time.Since(t1))

}
