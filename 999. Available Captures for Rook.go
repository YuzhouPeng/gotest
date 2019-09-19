package leetcode_go

func numRookCaptures(board [][]byte) int {
	var length int = len(board)
	var x int
	var y int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if string(board[i][j]) == "R" {
				x = i
				y = j
				break
			}
		}
	}
	var count int = 0
	for i := x - 1; i >= 0; i-- {
		if string(board[i][y]) != "." {
			if string(board[i][y]) == "p" {
				count += 1
				break
			} else if string(board[i][y]) == "B" {
				break
			}
		}
	}
	for i := x + 1; i < len(board); i++ {
		if string(board[i][y]) != "." {
			if string(board[i][y]) == "p" {
				count += 1
				break
			} else if string(board[i][y]) == "B" {
				break
			}
		}
	}
	for i := y - 1; i >= 0; i-- {
		if string(board[x][i]) != "." {
			if string(board[x][i]) == "p" {
				count += 1
				break
			} else if string(board[x][i]) == "B" {
				break
			}
		}
	}
	for i := y + 1; i < len(board); i++ {
		if string(board[x][i]) != "." {
			if string(board[x][i]) == "p" {
				count += 1
				break
			} else if string(board[x][i]) == "B" {
				break
			}
		}
	}
	return count
}
