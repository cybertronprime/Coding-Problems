package main

func HowSum(val int, param []int, myMap map[int]int) int {
	//Optimised
	if val == 0 {
		return 1
	}
	if val < 0 {
		return 0
	}
	if _, exists := myMap[val]; exists {
		return myMap[val]
	}
	counter:=0
	for _, x := range param {
		diff := val - x
		myMap[val] = HowSum(diff, param, myMap) 
		counter=counter+myMap[val]

	}
	return counter
}
