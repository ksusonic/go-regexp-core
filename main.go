package main

import (
	"fmt"
	"log"
	"strings"
)

const DEBUG = false

func scanInput() (string, string) {
	var fullInput string
	_, _ = fmt.Scanf("%s", &fullInput)
	split := strings.Split(fullInput, "|")
	if len(split) != 2 {
		log.Fatalln("Incorrect input! Expexted: a|b")
	}
	return split[0], split[1]
}

func CharMatch(regexpChar, char uint8) bool {
	return regexpChar == uint8('.') || regexpChar == char
}

func PatternMatch(regex, input string) bool {
	for i := range regex {
		if !CharMatch(regex[i], input[i]) {
			return false
		}
	}
	return true
}

func RegexMatch(regex, input string) bool {
	regEndPos := len(regex)
	for regEndPos <= len(input) {
		if PatternMatch(regex, input[regEndPos-len(regex):regEndPos]) {
			return true
		}
		regEndPos++
	}
	return false
}

func main() {
	regex, input := scanInput()

	if DEBUG {
		log.Printf("comparing '%s' and '%s'", regex, input)
	}

	if regex == "" {
		fmt.Println(true)
		return
	}

	fmt.Println(RegexMatch(regex, input))
}
