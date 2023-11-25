package main

func HowConstruct(target string, wordBank []string, myMap map[string]int) int {
	if target == "" {
		return 1

	}
	if _, exists := myMap[target]; exists {
		return myMap[target]
	}
	ifFound := 0
	for _, letter := range wordBank {
		letterLen := len(letter)

		if len(target) >= letterLen && target[0:letterLen] == letter {

			start := letterLen
			end := len(target)
			value := HowConstruct(target[start:end], wordBank, myMap)

			ifFound = ifFound + value
		}

	}
	myMap[target] = ifFound
	return ifFound
}

func HowConstructTab(target string, wordBank []string) int {

	newArr:=make([]int,len(target)+1)
	newArr[0]=1

	for i,_ :=range newArr{
		if newArr[i]>0{

			for _,word:=range wordBank{
				lenOfWord:=len(word)
				if lenOfWord+i<=len(target) && target[i:i+lenOfWord]==word{
					newArr[i+lenOfWord]+=newArr[i]
				}

			}

		}



	}
	return newArr[len(target)]


}

