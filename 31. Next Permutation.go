package leetcode_go

import "sort"

func nextPermutation(nums []int) {
	if len(nums) > 1 {
		var flag int = 0
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {
				flag = 1
				var temp int
				var min int = nums[i]
				var jadd int = i
				for j := i; j < len(nums); j++ {
					if min >= nums[j] && nums[j] > nums[i-1] {
						min = nums[j]
						jadd = j
					}
				}
				temp = nums[jadd]
				nums[jadd] = nums[i-1]
				nums[i-1] = temp

				reverse(nums, i)

				break
			}
		}
		if flag == 0 {
			sort.Ints((nums))
		}
	}
}

func reverse(nums []int, starti int) {
	var start int = starti
	var end int = len(nums) - 1
	for {
		if start >= end {
			break
		}
		var temp int
		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start += 1
		end -= 1
	}
}
