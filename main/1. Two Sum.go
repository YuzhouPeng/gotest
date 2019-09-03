package main

func twoSum(nums []int, target int) []int {
	var result []int
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i != j && (v1+v2) == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}
