package main

func CanSum(val int, param []int, myMap map[int]bool) bool {
	// //Unoptimised
	// if val == 0 {
	// 	return true
	// }
	// if val < 0 {
	// 	return false
	// }

	// for _, x := range param {

	// 	if Prog3(val-x, param) == true {
	// 		return true
	// 	}

	// }
	// return false

	//Optimised
	if val == 0 {
		return true
	}
	if val < 0 {
		return false
	}
	if _, exists := myMap[val]; exists {
		return myMap[val]
	}

	for _, x := range param {
		diff := val - x

		if CanSum(diff, param, myMap) == true {
			myMap[val] = true

			return true
		}
		myMap[val] = false

	}
	return false
}
