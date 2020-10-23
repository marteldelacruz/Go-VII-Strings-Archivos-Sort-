package main

import (
	"fmt"
	"os"
	"sort"
)

const (
	ASC = "asc"
	DES = "des"
)

func sortList(values *[]string, order string) {
	// sort values

	switch order {
	case ASC:
		sort.Strings(*values)
		break
	case DES:
		sort.Sort(sort.Reverse(sort.StringSlice(*values)))
		break
	}
}

func saveToFile(values *[]string, fileName string) bool {
	// creating file
	file, err := os.Create(fileName)

	// verify if the files has been created
	if err != nil {
		return false
	}

	// writing content
	for _, s := range *values {
		file.WriteString(s + "\n")
	}

	// close file
	defer file.Close()

	return true
}

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

	// saving asc file
	sortList(&list, ASC)
	if saveToFile(&list, "asecendente.txt") {
		fmt.Println("asecendente.txt file saved...")
	}

	// saving des file
	sortList(&list, DES)
	if saveToFile(&list, "descendente.txt") {
		fmt.Println("descendente.txt file saved...")
	}
}
