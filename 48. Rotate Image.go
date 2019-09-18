package leetcode_go

func rotate(matrix [][]int) {
	var length = len(matrix)
	var result []int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			result = append(result, matrix[i][j])
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j] = result[length*(length-j)+i-length]
		}
	}
}
