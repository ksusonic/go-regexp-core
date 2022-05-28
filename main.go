package main

import (
	"fmt"
	"log"
	"strings"
)

const DEBUG = false

func scanInput() (string, string) {
	var fullInput string
	fmt.Scanf("%s", &fullInput)
	split := strings.Split(fullInput, "|")
	if len(split) != 2 {
		log.Fatalln("Incorrect input! Expexted: a|b")
	}
	return split[0], split[1]
}

func main() {
	regex, input := scanInput()

	if DEBUG {
		log.Printf("comparing '%s' and '%s'", regex, input)
	}

	// pre-section of code
	if regex == input || regex == "" {
		fmt.Println(true)
		return
	}

	if regex == "." && len(input) == 1 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
