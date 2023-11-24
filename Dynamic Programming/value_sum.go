package main

import "fmt"

func ValueSum(val int, param []int, myMap map[int][]int) []int {
	//Optimised
	//tells about  the first nuumber that adds up to be the number given as input
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
		diff := val - x
		vals := ValueSum(diff, param, myMap)
		if vals != nil {

			vals = append(vals, x)
			myMap[val] = vals
			return vals
		}

	}
	myMap[val] = nil

	return nil
}

func ValueSumTab(val int, param []int) []int {

	newArr:=make([][]int,val+1)
	newArr[0]=[]int{0}

	for i,_:=range newArr{
		currentLength:=len(newArr[i])

		if  currentLength>0 {

			for _,data:=range param{

				if i+data<=val{
					
					oldArray:=newArr[i]
					oldArray=append(oldArray, data)
					newArr[i+data]=oldArray
			
					
				}
				
	
	
			}

		}
	}
	fmt.Println(newArr)
	return newArr[val][1:]







}
