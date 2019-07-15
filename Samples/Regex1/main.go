package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	/*
		EXEMPLO DE REGX SO NUMERO
	*/

	example := "ss216.399.218-77"

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(example, "")

	fmt.Printf("A string of %s becomes %s \n", example, processedString)
}
