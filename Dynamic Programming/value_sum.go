package main

func ValueSum(val int, param []int, myMap map[int]bool, vals []int) (val1 bool, val3 []int) {
	//Optimised
	//tells about  the first nuumber that adds up to be the number given as input
	if val == 0 {
		return true, vals
	}
	if val < 0 {
		return false, vals
	}
	if _, exists := myMap[val]; exists {
		return myMap[val], vals
	}
	for _, x := range param {
		diff := val - x
		xyz, vals := ValueSum(diff, param, myMap, vals)
		if xyz == true {
			myMap[diff] = true
			vals = append(vals, x)
			return true, vals

		}

		myMap[diff] = false

	}

	return false, vals
}



