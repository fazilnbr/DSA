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

func SelectionSort(arr []int, limit int) []int {
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

func QuickSort(arr []int, limit int) []int {
	quicksorthelper(arr, 0, limit-1)
	return arr
}
func quicksorthelper(arr []int, startIdx int, endIdx int) {
	if startIdx >= endIdx {
		return
	}
	pivotIdx := startIdx
	leftIdx := startIdx + 1
	rightIdx := endIdx

	for leftIdx <= rightIdx {
		if arr[pivotIdx] < arr[leftIdx] && arr[pivotIdx] > arr[rightIdx] {

			arr[leftIdx], arr[rightIdx] = arr[rightIdx], arr[leftIdx]

		}
		if arr[leftIdx] <= arr[pivotIdx] {
			leftIdx++
		}
		if arr[rightIdx] >= arr[pivotIdx] {
			rightIdx--
		}
	}
	arr[rightIdx], arr[pivotIdx] = arr[pivotIdx], arr[rightIdx]

	quicksorthelper(arr, startIdx, rightIdx-1)
	quicksorthelper(arr, rightIdx+1, endIdx)

}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
