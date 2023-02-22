package main

func CanConstruct(target string, wordBank []string) bool {
	if target == "" {
		return true

	}

	for _, letter := range wordBank {
		if target != "" && letter == string(target[0]) {

			if len(target) > 1 {
				ifFound := CanConstruct(target[1:len(target)-1], wordBank)
				if ifFound == true {
					return ifFound
				}

			} else {
				ifFound := CanConstruct("", wordBank)
				if ifFound == true {
					return ifFound
				}
			}

		}

	}

	return false

}
