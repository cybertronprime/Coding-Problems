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
