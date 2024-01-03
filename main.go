package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func trimPrefixAndSuffix(regex string) string {
	regex = strings.TrimPrefix(regex, "^")
	regex = strings.TrimSuffix(regex, "$")
	return regex
}

func lenghtCheck(elements []string) bool {
	return len(elements) == 0
}
func metaMatching(regex, input string) bool {
	var result bool
	metaCache := []rune{'^', '$', '.', '?', '*', '+'}

	for _, r := range regex {
		for _, m := range metaCache {
			if strings.IndexRune(regex, m)))
		})
	}
	regex = trimPrefixAndSuffix(regex)

	// ^ stars with
	// $ ends with
	// . matches with
	// * preceding char zero or more
	// ? preceding char zero or one
	// + preceding char one or more
	switch result {
	case len(regex) == len(input):

	}
	// if len(regex) == len(input) {
	//
	return result
}

func matching(regex, input string) bool {

	var result bool
	inputElements := strings.Fields(input)

	if strings.Contains(input, regex) {
		result = true
		return result
	}

	if lenghtCheck(inputElements) {
		result = false
		return result
	}

	if len(inputElements) == 1 && !strings.HasPrefix(regex, "^") && !strings.HasSuffix(regex, "$") {
		result = metaMatching(regex, input)
		return result
	}

	if strings.HasPrefix(regex, "^") && strings.HasSuffix(regex, "$") {
		regex = trimPrefixAndSuffix(regex)
		if len(inputElements) > 1 {
			result = false
			return result
		}
		result = (regex == inputElements[0])
		return result

	}
	if strings.HasPrefix(regex, "^") {
		regex = trimPrefixAndSuffix(regex)
		if regex == "." {
			result = true
			return result
		}
		result = strings.HasPrefix(inputElements[0], regex)
		return result
	}
	if strings.HasSuffix(regex, "$") {
		regex = trimPrefixAndSuffix(regex)
		if regex == "." {
			result = true
			return result
		}
		result = strings.HasSuffix(inputElements[len(inputElements)-1], regex)
		return result
	}
	return result
}

func extractStirng(part string) (string, error) {
	if part == "" {
		return "err", fmt.Errorf("this part is empty")
	}
	return part, nil
}

func frmStdin() bool {
	var input string

	reader := bufio.NewReader(os.Stdin)
	defer reader.Reset(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(input, "|") {
		return false
	}
	parts := strings.SplitN(input, "|", 2)
	a, err := extractStirng(parts[0])
	if err != nil {
		return true
	}

	b, err := extractStirng(parts[1])
	if err != nil {
		return false
	}
	return matching(string(a), string(b))
}

func main() {

	fmt.Printf("%v\n", frmStdin())
}
