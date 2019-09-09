package leetcode_go

func rotate(nums []int, k int) {
	newarray := []int{}
	if len(nums) > 1 {
		if len(nums) > k {
			for i := 0; i < k; i++ {
				newarray = append(newarray, nums[len(nums)+i-k])
			}
			for i := 0; i < len(nums)-k; i++ {
				newarray = append(newarray, nums[i])
			}
			copy(nums, newarray)
		} else {
			var newk int = k % len(nums)
			for i := 0; i < newk; i++ {
				newarray = append(newarray, nums[len(nums)+i-newk])
			}
			for i := 0; i < len(nums)-newk; i++ {
				newarray = append(newarray, nums[i])
			}
			copy(nums, newarray)
		}
	}

}
