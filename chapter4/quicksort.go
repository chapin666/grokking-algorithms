package chapter4

func quicksortAdvance(arr []int, left, right int) {
	if left < right {
		pos := partition(arr, left, right)
		quicksortAdvance(arr, left, pos-1)
		quicksortAdvance(arr, pos+1, right)
	}
}

func partition(arr []int, left, right int) int {
	key := arr[right]
	i := left

	for j := left; j < right; j++ {
		if arr[j] < key {
			// 小于key的移动到左边
			swap(arr, i, j)
			i++
		}
	}

	// 移动最右边的元素
	swap(arr, i, right)

	return i
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func quicksort(arr []int, left, right int) {
	if left < right {
		pos := (left + right) / 2
		key := arr[pos]
		i := left
		j := right

		for {
			// left
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i >= j {
				break
			}
			swap(arr, i, j)
		}

		quicksort(arr, left, i-1)
		quicksort(arr, j+1, right)
	}
}
