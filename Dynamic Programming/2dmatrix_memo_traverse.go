package main

import (


	"strconv")

func TwoDmatrixTraverse(rows int, column int, myMap map[string]int) int {

	if rows<=0 || column<=0{
		return 0
	}
	if rows==1 && column==1{
		return 1
	} 
	xnew:=strconv.Itoa(rows)
	xnew=xnew+strconv.Itoa(column)
	if _,exists:=myMap[xnew]; exists{
		return myMap[xnew]
	}
	leftSide:=	TwoDmatrixTraverse(rows-1,column,myMap)
	rightSide:=	TwoDmatrixTraverse(rows,column-1,myMap)
	myMap[xnew]=leftSide+rightSide

	return leftSide+rightSide
}


