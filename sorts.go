package main

func bubbleSort(arr []uint16) []uint16 {
	// swapped := true
	// m := 0
	// for swapped {
	// 	swapped = false
	// 	k := 0
	// 	for i := 1; i < len(arr)-m; i++ {
	// 		if arr[i] < arr[k] {
	// 			_k := arr[k]
	// 			_i := arr[i]
	// 			arr[i] = _k
	// 			arr[k] = _i
	// 			swapped = true
	// 		}
	// 		k = i
	// 	}
	// 	m++
	// }
	// return arr
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func quicksort(arr []uint16, low uint16, high uint16) []uint16 {
	if low < high {
		arr, p := _partition(arr, low, high)
		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)
	}
	return arr
}

func _partition(arr []uint16, low uint16, high uint16) ([]uint16, uint16) {
	p := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
