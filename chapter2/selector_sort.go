package chapter2

func selectorSort(arr []int) (newArray []int) {
	count := len(arr)
	for i := 0; i < count; i++ {
		index, _ := findMin(arr)
		newArray = append(newArray, arr[index])
		arr = append(arr[:index], arr[index+1:]...)
	}
	return newArray
}
func findMin(arr []int) (minIndex int, minNumber int) {
	minIndex = 0
	minNumber = arr[0]
	for index, value := range arr {
		if value < minNumber {
			minNumber = value
			minIndex = index
		}
	}

	return
}
