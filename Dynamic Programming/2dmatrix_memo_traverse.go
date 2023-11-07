package main

import (
	"fmt"
	"strconv"
)

func TwoDmatrixTraverse(rows int, column int, myMap map[string]int) int {

	if rows == 0 || column == 0 {
		return 0
	}
	if rows == 1 && column == 1 {
		return 1
	}
	xnew := strconv.Itoa(rows)
	xnew = xnew + strconv.Itoa(column)
	if _, exists := myMap[xnew]; exists {
		return myMap[xnew]
	}
	leftSide := TwoDmatrixTraverse(rows-1, column, myMap)
	rightSide := TwoDmatrixTraverse(rows, column-1, myMap)
	myMap[xnew] = leftSide + rightSide

	return leftSide + rightSide
}
func TwoDmatrixTab(row, column int) int {

	rowaArr := make([][]int, row+1)

	for i, _ := range rowaArr {
		rowaArr[i] = make([]int, column+1)

	}

	rowaArr[1][1] = 1
	for i, innerArray := range rowaArr {

		for j, _ := range innerArray {
			if i+1 <= row {
				rowaArr[i+1][j] += innerArray[j]

			}
			if j+1 <= column {
				rowaArr[i][j+1] += innerArray[j]

			}

		}

	}
	for i, _ := range rowaArr {
		fmt.Println(rowaArr[i])

	}

	return rowaArr[row][column]

}
