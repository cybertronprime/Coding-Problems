package main

func IndexCheck(array1 []int, target int) []int {
	resultArr := []int{-1, -1}
	currentIndex := 0

	for i, val := range array1 {
		if val == target && currentIndex > 0 {
			resultArr[1] = i
		}
		if val == target && currentIndex == 0 {
			resultArr[currentIndex] = i
			currentIndex++
			//Could also check if the next position is same as the current number , if not this is a [i,-1]
			//condition - so could break here
		}
	}
	if resultArr[1] == -1 {
		resultArr[1] = resultArr[0]
	}
	return resultArr
}
