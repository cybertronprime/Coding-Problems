package main

func Anagrams(string1, string2 string) bool {
	returnVal := false
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	if len(string1) != len(string2) {
		return returnVal
	}

	for i, _ := range string1 {
		letter := string(string1[i])
		if _, exists := map1[letter]; exists {
			map1[letter] = map1[letter] + 1
		} else {
			map1[letter] = 1
		}

	}
	for i, _ := range string2 {
		letter := string(string2[i])
		if _, exists := map2[letter]; exists {
			map2[letter] = map2[letter] + 1
		} else {
			map2[letter] = 1
		}

	}
	if len(map1) != len(map2) {
		return returnVal
	}
	for key, element := range map1 {

		if _, exists := map2[key]; exists {
			if map2[key] == element {
				returnVal = true
				continue
			} else {
				break
			}
		} else {
			break

		}

	}

	return returnVal

}
