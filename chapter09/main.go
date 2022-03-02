package main

import "fmt"

func main() {
	var title string
	var copies int
	var edition string
	title = "For the Love of Go"
	copies = 99
	edition = "first"
	fmt.Println(title)
	fmt.Println(copies)
	fmt.Println(edition)

	printTitle(title)
}

func printTitle(title string) {
	fmt.Println("Title: ", title)
}