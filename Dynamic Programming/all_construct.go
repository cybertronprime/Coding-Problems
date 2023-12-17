package main

import "fmt"

func AllConstruct(target string, wordBank []string, myMap map[string][][]string) [][]string {

	if target == "" {
		return [][]string{}
	}
	if _, ok := myMap[target]; ok {
		return myMap[target]
	}
	data := make([][]string, 0)

	for _, letter := range wordBank {

		if len(target) >= len(letter) && target[:len(letter)] == letter {

			value := AllConstruct(target[len(letter):], wordBank, myMap)

			if value != nil {
				if len(value) == 0 {
					letter1 := []string{letter}
					value = append(value, letter1)
				} else {
					for i, _ := range value {
						value[i] = append(value[i], letter)
					}

				}
				data = append(data, value...)

			}

		}
	}
	// myMap[target] = make([][]string, len(data))
	copy(myMap[target], data)
	fmt.Println(data, myMap)
//If i try to return myMap[target], the value of the map gets updated due to future operations, but when i return data, 
//its not affecting the map as map has a copy of data. and for future operations the data gets updated. 
//this all happens since they are passsed by refn.
	return data
}

func AllConstructTab(target string, wordBank []string)  [][]string {

	myArr:=make([][][]string,len(target)+1)

	newArr:=[][]string{}
	data:=""
	sencondArr:=[]string{data}
	newArr = append(newArr,sencondArr )

	myArr[0]=newArr

	for i:=range myArr{

		if len(myArr[i])>0{

			for _,word:=range wordBank{

				if i+len(word)<=len(target) && target[i:i+len(word)]==word{
					if len(myArr[i])==0{
						data1:=[]string{word}
						myArr[i+len(word)]=append(myArr[i+len(word)], data1)

					}else{
						for _,newData:=range myArr[i]{
							newData=append(newData,word)
							myArr[i+len(word)]=append(myArr[i+len(word)], newData)

						}

					}


				}



			}

		}
	}


fmt.Println(myArr)

return myArr[len(target)]



}