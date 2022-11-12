package sort

func BubbleSort(arr []int, limit int) []int {
	for i := 0; i < limit; i++ {
		for i := 1; i < limit; i++ {
			if arr[i] < arr[i-1] {
				temp := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = temp
			}
		}
	}
	return arr
}

func InsertionSort(arr []int, limit int) []int {
	for i := 1; i < limit; i++ {
		t := i - 1
		for t > 0 && arr[i] < arr[t] {
			temp := arr[i]
			arr[i] = arr[t]
			arr[t] = temp
		}
	}
	return arr
}

func SelectonSort(arr []int, limit int) []int {
	var min int
	for i := 0; i < limit; i++ {
		min = i
		for j := i; j < limit; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	return arr
}
