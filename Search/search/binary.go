package search

import "time"

func BainarySearch(target int, arr []int) (bool, int) {
	time.Sleep(1000)
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := start + ((end - start) / 2)
		if arr[mid] == target {
			return true, mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false, -1

}
