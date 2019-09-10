package leetcode_go

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for {
		if start >= end {
			break
		}
		var temp int = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start += 1
		end -= 1
	}
}
