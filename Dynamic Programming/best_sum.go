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

			if mapdata, exists := myMap[val]; !exists || len(mapdata) > len(val3) {
				myMap[val] = val3

			}

		}
		fmt.Println("myMap[val]", val, myMap[val])
	}
	// myMap[val] = nil

	return myMap[val]
}

func BestSumTab(val int, param []int) []int {

	newArr:=make([][]int,val+1)
	newArr[0]=[]int{0}

	for i,_:=range newArr{
		currentLength:=len(newArr[i])

		if  currentLength>0 {

			for _,data:=range param{

				if i+data<=val{
					
					oldArray:=newArr[i]
					oldArray=append(oldArray, data)
					existingData:=newArr[i+data]
					if len(existingData)==0 || len(existingData)>len(oldArray){
						newArr[i+data]=oldArray

					}
					
			
					
				}
				
	
	
			}

		}
	}
	fmt.Println(newArr)
	return newArr[val][1:]







}
