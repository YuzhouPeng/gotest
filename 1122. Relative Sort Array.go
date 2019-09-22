package leetcode_go

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	map1 := make(map[int]int)
	for i := 0; i < len(arr2); i++ {
		map1[arr2[i]] = 0
	}
	var remain []int
	for i := 0; i < len(arr1); i++ {
		var v int
		var ok bool
		v, ok = map1[arr1[i]]
		if !ok && v == v {
			remain = append(remain, arr1[i])
		} else {
			map1[arr1[i]] += 1
		}
	}
	var result []int
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < map1[arr2[i]]; j++ {
			result = append(result, arr2[i])
		}
	}
	sort.Ints(remain)
	for i := 0; i < len(remain); i++ {
		result = append(result, remain[i])
	}
	return result
}
