package leetcode_go

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := strings.ToLower(reg.ReplaceAllString(s, ""))
	fmt.Println(processedString)
	for i := 0; i < len(processedString); i++ {
		var j int = len(processedString) - 1 - i
		if i < j && processedString[i] != processedString[j] {
			return false
		}
	}
	return true
}
