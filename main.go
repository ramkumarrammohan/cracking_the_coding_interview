package main

import (
	"fmt"

	chapter1 "github.com/ramkumarrammohan/cracking_coding_interview/chapter-1"
)

func routeToChapter(chapterNo int) {
	switch chapterNo {
	case 1:
		chapter1.Chapter1Index()
	default:
		fmt.Println("Unknown chapter entered")
	}
}

func main() {
	fmt.Println("Welcome to Cracking Coding interview solutions...")
	fmt.Printf("Enter the chapter no: ")
	var chapter int
	fmt.Scanln(&chapter)
	routeToChapter(chapter)
}
