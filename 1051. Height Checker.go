package leetcode_go

import (
	"sort"
)

func heightChecker(heights []int) int {
	var cmp []int
	var count int = 0
	for i := 0; i < len(heights); i++ {
		cmp = append(cmp, heights[i])
	}
	sort.Ints(cmp)
	for i := 0; i < len(heights); i++ {
		if cmp[i] != heights[i] {
			count += 1
		}
	}
	return count
}
