package main

import "fmt"

func BestSum(val int, param []int, myMap map[int][]int) (val3 []int) {
	//Optimised
	//This method can be used in howsum
	//Need to understand how a different value gets returned at the end .
	if val == 0 {
		return []int{}
	}
	if val < 0 {
		return nil
	}
	if _, exists := myMap[val]; exists {
		return myMap[val]
	}
	for _, x := range param {

		val3 := BestSum(val-x, param, myMap)
		if val3 != nil {
			val3 = append(val3, x)
			if _, exists := myMap[val]; exists {

				if (len(val3) <= len(myMap[val])) {
					myMap[val] = val3
				}else{
					val3=myMap[val]
				}
			} else {
				myMap[val] = val3

			}

			fmt.Println("myMap[val]", val, myMap[val])

		}


	}
	return myMap[val]
}

