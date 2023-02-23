package main

import "fmt"

func CanConstruct(target string, wordBank []string) bool {
	if target == "" {
		return true

	}

	for _, letter := range wordBank {
		fmt.Println("target", target)
		fmt.Println("letter", letter)
		letterLen := len(letter)

		if len(target) >= letterLen && target[0:letterLen] == letter {

			start := letterLen
			end := len(target)

			if len(target) == letterLen {
				start = 0
				end = 0

			}
			ifFound := CanConstruct(target[start:end], wordBank)
			if ifFound == true {
				return ifFound
			}

		}

	}

	return false

}
