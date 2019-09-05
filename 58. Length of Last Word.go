package leetcode_go

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	news := strings.Split(s, " ")
	return len(string(news[len(news)-1]))
}
