package main

func AllConstruct(target string, wordBank []string) [][]string {

	if target == "" {
		return [][]string{}
	}

	for _, letter := range wordBank {

		if len(target) >= len(letter) && target[:len(letter)] == letter {


			value:=AllConstruct(target[len(letter):],wordBank)
			if value != nil{



			}



		}
	}

	return nil
}
