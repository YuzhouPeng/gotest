package leetcode_go

func rotate(nums []int, k int) {
	k = k % len(nums)
	var count int = 0
	for start := 0; count < len(nums); start++ {
		var current int = start
		var prev int = nums[start]
		for {
			var next int = (current + k) % len(nums)
			var temp int = nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count += 1
			if start == current {
				break
			}
		}
	}

}
