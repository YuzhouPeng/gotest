package leetcode_go

func getRow(rowIndex int) []int {
	var result []int

	for i := 0; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = result[j] + result[j-1]
		}
		result = row
	}
	return result
}
