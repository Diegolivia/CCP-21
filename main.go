package main

import "fmt"

func MakeBoard(reinas [8]int) {
	var board [8][8]int
	for i := 0; i < 8; i++ {
		board[i] = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	}
	for idx, elem := range reinas {
		if elem != -1 {
			board[idx][elem] = 1
		}
	}
	for _, val := range board {
		fmt.Println(val)
	}
}

func contains(arr [3]int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

func Validate(reinas [8]int, row int, col int) bool {
	delta := 0
	col_i := 0
	for row_i := 0; row_i < row; row_i++ {
		col_i = reinas[row_i]
		delta = row - row_i
		arr := [3]int{col_i, col_i + delta, col_i - delta}
		if contains(arr, col) {
			return false
		}
	}
	return true
}

func nreinas(reinas [8]int, row int) {
	n := len(reinas)
	if row == n {
		MakeBoard(reinas)
		println("Done")
	} else {
		for col := 0; col < n; col++ {
			if Validate(reinas, row, col) {
				reinas[row] = col
				nreinas(reinas, row+1)
			}
		}
	}
}

func main() {
	reina := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}
	nreinas(reina, 0)
}
