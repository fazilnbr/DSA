package search

import "time"

func LeaniarSearch(target int, arr []int) (bool, int) {
	time.Sleep(1000)
	for i, val := range arr {
		if val == target {
			return true, i
		}
	}
	return false, -1

}
