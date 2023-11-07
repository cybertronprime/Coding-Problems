package main

import (
	"fmt"
	"math"
)

func BinarySearch(arr1 []int, target int) int {

	low := 0
	high := len(arr1) - 1
	for high > low {
		fmt.Println(high, low)
		mid := int(math.Ceil((float64(high) + float64(low)) / 2))
		fmt.Println(mid, arr1[mid])
		if arr1[mid] == target {
			return mid
		} else if arr1[mid] < target {
			low = mid + 1
		} else if arr1[mid] > target {
			high = mid - 1
		}

	}
	return -1

}
