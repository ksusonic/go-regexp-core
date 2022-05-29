package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const DEBUG = false

func scanInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	split := strings.Split(scanner.Text(), "|")
	if len(split) != 2 {
		log.Fatalln("Incorrect input! Expexted: a|b")
	}
	return split[0], split[1]
}

func CharMatch(regexpChar, char uint8) bool {
	return regexpChar == uint8('.') || regexpChar == char
}

func PatternMatch(regex, input string) bool {
	if DEBUG {
		log.Printf("comparing regex '%s' to '%s'", regex, input)
	}

	for i := range regex {
		if !CharMatch(regex[i], input[i]) {
			return false
		}
	}
	return true
}

func ProcessRegexp(regex string) (string, bool, bool) {
	var startPref, endPref = strings.HasPrefix(regex, "^"), strings.HasSuffix(regex, "$")
	if startPref {
		regex = strings.TrimPrefix(regex, "^")
	}
	if endPref {
		regex = strings.TrimSuffix(regex, "$")
	}
	return regex, startPref, endPref
}

func RegexMatch(regex, input string) bool {
	regex, startFlag, endFlag := ProcessRegexp(regex)
	regEndPos := len(regex)

	if startFlag && endFlag {
		if len(regex) != len(input) {
			return false
		} else {
			return PatternMatch(regex, input)
		}
	} else if startFlag || endFlag {
		if len(regex) <= len(input) {
			if startFlag {
				return PatternMatch(regex, input[:len(regex)])
			} else /* endFlag */ {
				return PatternMatch(regex, input[len(input)-len(regex):])
			}
		} else {
			return false
		}
	} else {
		for regEndPos <= len(input) {
			if PatternMatch(regex, input[regEndPos-len(regex):regEndPos]) {
				return true
			}
			regEndPos++
		}
		return false
	}
}

func main() {
	regex, input := scanInput()

	if regex == "" {
		fmt.Println(true)
		return
	}

	fmt.Println(RegexMatch(regex, input))
}
