package main

func Fibb(value int, myMap map[int]int) int {
	if value <= 2 {
		return 1
	}
	if _, ok := myMap[value]; ok {
		return myMap[value]
	}
	newValue := Fibb(value-1, myMap) + Fibb(value-2, myMap)
	myMap[value] = newValue
	return newValue

}

func FibbTab(value int) int {

	array := make([]int, value+1)
	array[0] = 0
	array[1] = 1

	for i, _ := range array {

		if i+2 == value {
			array[i+2] += array[i]
		}
		if i+1 == value {
			array[i+1] += array[i]
		}

	}
	return array[value]

}
