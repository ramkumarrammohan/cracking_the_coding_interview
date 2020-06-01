package chapter1

import "fmt"

// Chapter1Index Routes to chapter-1 problems
func Chapter1Index() {
	fmt.Printf("Enter the problem no: ")
	var problemNo int
	fmt.Scanln(&problemNo)

	switch problemNo {
	case 1:
		isUnique()
	default:
		fmt.Println("Unknow problem chosen. Try another")
	}
}
