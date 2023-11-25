package main

import (
	"fmt"
)

func main() {

	//----------------------------------Program 1 -----------------------------------// FIBBONACI
	//For FIbbonaci
	// map1 := make(map[int]int)
	// fmt.Println(Fibb(50, map1))
	//----------------------------------Program 2 -----------------------------------// GRID TRAVELLER
	//For 2D Matrix Traverse

	// type matrixKey struct{
	// 	one int
	// 	two int

	// }
	// map1 := make(map[matrixKey]int)
	// fmt.Println(twoDmatrixTraverse(2,3,matrixKey, map1))

	// map1 := make(map[string]int)
	// fmt.Println(TwoDmatrixTraverse(1000,40000005, map1))
	//----------------------------------Program 3 -----------------------------------// CAN SUM
	//Returning true if numbers from the array add up to be the number given as input
	// m := make(map[int]bool)
	// var newarr= []int{7,14}
	// fmt.Println(CanSum(10000000, newarr,m))


	// //----------------------------------Program 4 -----------------------------------// HOW SUM
	// 	//Returning numbers of occurences if numbers from the array add up to be the number given as input
	// m := make(map[int]int)
	// var newarr= []int{1,2,3}
	// fmt.Println(HowSum(3, newarr,m))

	//----------------------------------Program 5 -----------------------------------//  VALUE SUM
	//Returning numbers  if numbers from the array add up to be the number given as input
	// m := make(map[int][]int)
	// var newarr = []int{1, 2, 3}

	// fmt.Println(ValueSum(3, newarr, m))

	//----------------------------------Program 6 -----------------------------------// BEST SUM
	//Returning least numbers  if numbers from the array add up to be the number given as input
	// m := make(map[int][]int)

	// var newarr= []int{2,8}
	// fmt.Println(BestSum(12, newarr,m))

	//----------------------------------Program 7 -----------------------------------//
	//Returning if target word can be formed by the given wordbank

	// var newarr = []string{"e","ee","eee","eeeeee","eeeeeeeeee"}
	// var myMap =make(map[string]bool)
	// fmt.Println(CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", newarr,myMap))

	//----------------------------------Program 8 -----------------------------------//
	//Returning numbers of times a target word can be formed by the given wordbank

	// var newarr = []string{"ab","abc","cd","def","abcd"}
	// var myMap =make(map[string]int)
	// fmt.Println(HowConstruct("abcdef", newarr,myMap))
	//----------------------------------Program 8 -----------------------------------// ALL CONSTRUCT
	// var newarr = []string{"ab", "abc", "cd", "ef", "abcd"}

	// fmt.Println(FibbTab(50))
	//----------------------------------Program 9 -----------------------------------// FIB TABULATION

	// fmt.Println(TwoDmatrixTab(3,3))


	//----------------------------------Program 10-----------------------------------// CAN_SUM TABULATION
	// var newarr= []int{5,3,4}
	// fmt.Println(CanSumTab(7, newarr))
		//----------------------------------Program 11-----------------------------------// HOW_SUM TABULATION
	// var newarr= []int{5,3,4}
	// fmt.Println(HowSumTab(10, newarr))

		//----------------------------------Program 12-----------------------------------// Value_SUM TABULATION

	// var newarr= []int{5,3,4}
	// fmt.Println(ValueSumTab(7, newarr))
			//----------------------------------Program 13-----------------------------------// Best_SUM TABULATION

			// var newarr= []int{5,3,4,}
			// fmt.Println(BestSumTab(7, newarr))


//----------------------------------Program 14-----------------------------------// CanConstructTab TABULATION

	// var newarr = []string{"ab","abc","cd","def","abcd"}
	// fmt.Println(CanConstructTab("abcdef", newarr))
		//----------------------------------Program 15-----------------------------------//
	//Returning numbers of times a target word can be formed by the given wordbank
	var newarr = []string{"ab","abc","cd","def","abcd","abcdef"}
	fmt.Println(HowConstructTab("abcdef", newarr))
}
