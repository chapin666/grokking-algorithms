package chapter1

import (
	"log"
)

// 递归方式实现二分查找
func binarySearchRecursion(datas interface{}, start, end, key int) int {
	arr, ok := datas.([]int)
	if ok {
		mid := (end-start)/2 + start
		if arr[mid] == key {
			return mid
		}
		if start >= end {
			return -1
		} else if key > arr[mid] {
			return binarySearchRecursion(arr, mid+1, end, key)
		} else if key < arr[mid] {
			return binarySearchRecursion(arr, start, mid-1, key)

		}
	} else {
		log.Fatalln("the parameter 'data' Must be an array")
	}
	return -1
}

// 二分顺序查找
func binarySearchOrder(datas interface{}, start, end, key int) (position int) {
	arr, ok := datas.([]int)
	if ok {
		for start < end {
			mid := (end-start)/2 + start
			if arr[mid] > key {
				end = mid - 1
			} else if arr[mid] < key {
				start = mid + 1
			} else {
				return mid
			}
		}
	} else {
		log.Fatalln("the parameter 'data' Must be an array")
	}
	return -1
}
