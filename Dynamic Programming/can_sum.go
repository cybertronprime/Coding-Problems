package main

func CanSum(val int, param []int, myMap map[int]bool) bool {

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
		

	}
	myMap[val] = false
	return false
}

func CanSumTab(val int, param []int) bool {

	newArr:=make([]bool,val+1)
	newArr[0]=true

	for i,_:=range newArr{

		if  newArr[i]==true {

			for _,data:=range param{

				if i+data<=val{
	
					newArr[data+i]=true
				}
	
	
			}

		}
		

	}
	return newArr[val]

}

