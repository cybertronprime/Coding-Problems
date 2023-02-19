package main

import (
	"fmt"
	
)


func main() {
 
	//----------------------------------Program 1 -----------------------------------//
	//For FIbbonaci
	// map1 := make(map[int]int)
	// fmt.Println(Fibb(50, map1))
	//----------------------------------Program 2 -----------------------------------//
	//For 2D Matrix Traverse

	// type matrixKey struct{
	// 	one int
	// 	two int

	// }
	// map1 := make(map[matrixKey]int)
	// fmt.Println(twoDmatrixTraverse(2,3,matrixKey, map1))

	// map1 := make(map[string]int)
	// fmt.Println(TwoDmatrixTraverse(1000,40000005, map1))
	//----------------------------------Program 3 -----------------------------------//
	//Returning true if numbers from the array add up to be the number given as input
	// m := make(map[int]bool)
	// var newarr= []int{7,14}
	// fmt.Println(CanSum(10000000, newarr,m))


// //----------------------------------Program 4 -----------------------------------//
// 	//Returning numbers of occurences if numbers from the array add up to be the number given as input
// 	m := make(map[int]int)
// 	var newarr= []int{5,3,4,7}
// 	fmt.Println(HowSum(7, newarr,m))

//----------------------------------Program 5 -----------------------------------//
	//Returning numbers  if numbers from the array add up to be the number given as input
	// m := make(map[int]bool)
	// var newarr= []int{7,5,4,3}
	// var newarr1= []int{}
	// fmt.Println(ValueSum(7, newarr,m,newarr1))


//----------------------------------Program 6 -----------------------------------//
	//Returning least numbers  if numbers from the array add up to be the number given as input
	m := make(map[int][]int)
	var newarr= []int{2,3,5,7}
	fmt.Println(BestSum(7, newarr,m))

}

