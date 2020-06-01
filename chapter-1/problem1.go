package chapter1

import "fmt"

func isUnique() {
	fmt.Printf("Enter a string: ")
	var enteredString string
	fmt.Scanln(&enteredString)

	ret := runAlgorithm(enteredString)
	if ret {
		fmt.Println("Unique characters")
	} else {
		fmt.Println("Non unique characters")
	}
}

func runAlgorithm(value string) bool {
	if len(value) > 26 {
		return false
	}

	holder := [26]bool{}

	for _, char := range value {
		var index int = -1
		if char >= 'a' && char <= 'z' {
			index = int(char) - int('a')
		} else if char >= 'A' && char <= 'Z' {
			index = int(char) - int('A')
		} else {
			fmt.Println("only accpting alpha chars in string")
			return false
		}

		if holder[index] {
			return false
		}
		holder[index] = true
	}

	return true
}
