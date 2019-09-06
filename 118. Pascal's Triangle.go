package leetcode_go

func generate(numRows int) [][]int {
	var newarray [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = newarray[i-1][j-1] + newarray[i-1][j]
		}
		newarray = append(newarray, row)
	}
	return newarray

}
