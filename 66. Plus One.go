package leetcode_go

func plusOne(digits []int) []int {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	var i int = 0
	var flag int = 0
	for {
		if i == 0 {
			digits[i] += 1
			if digits[i] > 9 {
				digits[i] = 0
				flag = 1
				if i == len(digits)-1 {
					digits = append(digits, 1)
					flag = 0
				}
			} else {
				break
			}
		}
		if i > len(digits)-1 || flag == 0 {
			break
		}
		if i != 0 {
			digits[i] += 1
			if digits[i] > 9 {
				digits[i] = 0
				if i == len(digits)-1 {
					digits = append(digits, 1)
					flag = 0
				}
			} else {
				flag = 0
			}
		}
		i += 1
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
