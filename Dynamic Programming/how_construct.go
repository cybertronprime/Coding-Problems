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
