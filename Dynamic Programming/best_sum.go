package main

func BestSum(val int, param []int, myMap map[int][]int, vals []int) (val3 []int) {
	//Optimised
	//This method can be used in howsum

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

		vals = BestSum(val-x, param, myMap, vals)
		if vals != nil {

			if _, exists := myMap[val]; exists {
				lengthVal := len(myMap[val])
				if len(vals) <= lengthVal {
					myMap[val] = append(vals, x)

				}
				continue

			}

			myMap[val] = append(vals, x)

		}

		myMap[val] = nil

	}
	return myMap[val]
}
