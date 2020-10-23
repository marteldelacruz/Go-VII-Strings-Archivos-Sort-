package main

import "fmt"

func main() {
	var amoutOfStrings int
	var list []string

	// get user input
	fmt.Print("Amount of strings: ")
	fmt.Scanln(&amoutOfStrings)

	for i := 0; i < amoutOfStrings; i++ {
		fmt.Printf("#%d: ", i+1)
		temp := ""
		// read new string
		fmt.Scanln(&temp)

		// save string on list
		list = append(list, temp)
	}

	// print list
	for _, s := range list {
		fmt.Println(s)
	}
}
