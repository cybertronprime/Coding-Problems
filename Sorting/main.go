package main

import "fmt"

func main() {

	//Problem 1 : Finding first and last occurence of target in a integer array: Sorted Array
	array1 := []int{1, 5, 6, 7, 7, 7, 8, 9, 9, 9, 9}
	target := 5

	// fmt.Println(IndexCheck(array1, target))
	//Problem 2 : Binary search
	fmt.Println(BinarySearch(array1, target))

}
