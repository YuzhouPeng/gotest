package leetcode_go

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := "1"
	n -= 1
	for {
		if n <= 0 {
			break
		}
		var j int = 0
		temp := ""
		for i := 0; i < len(s); i++ {
			if i+1 == len(s) || s[i] != s[i+1] {
				temp += strconv.Itoa(i-j+1) + string(s[i])
				j = i + 1
			}
		}
		s = temp
		n -= 1
	}
	return s

}
