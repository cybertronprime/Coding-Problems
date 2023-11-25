package main

func CanConstruct(target string, wordBank []string, myMap map[string]bool) bool {
	if target == "" {
		return true

	}
	if _, exists := myMap[target]; exists {
		return myMap[target]
	}

	for _, letter := range wordBank {
		letterLen := len(letter)

		if len(target) >= letterLen && target[0:letterLen] == letter {

			start := letterLen
			end := len(target)
			ifFound := CanConstruct(target[start:end], wordBank, myMap)
			if ifFound == true {
				myMap[target] = true

				return ifFound
			}

		}

	}
	myMap[target] = false
	return false

}
func CanConstructTab(target string, wordBank []string) bool {

	newArr:=make([]bool,len(target)+1)
	newArr[0]=true

	for i,_ :=range newArr{
		if newArr[i]{

			for _,word:=range wordBank{
				lenOfWord:=len(word)
				if lenOfWord+i<=len(target) && target[i:i+lenOfWord]==word{
					newArr[i+lenOfWord]=true
				}

			}

		}



	}
	return newArr[len(target)]


}

