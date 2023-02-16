package main
import (

	
)


func Prog3(val int , param []int) bool{

if val ==0 {
	return true
}
if val<0{
	return false
}

var value bool
for _, x := range param{

	value= Prog3(val-x,param)
	if value{
		break
	}
}


return  value
}