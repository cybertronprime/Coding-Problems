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
