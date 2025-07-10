package main

import (
	"fmt"
	"modularGO/FileHandling"
)

func main() {
	a := FileHandling.GetUserInput("Enter filename: ")
	b := FileHandling.GetUserInput("Write Content: ")

	err := FileHandling.SaveToFile(a, b)
	if err != nil {
		fmt.Println("Error saving file:", err)
	} else {
		fmt.Println("File saved successfully as:", a)
	}
}
