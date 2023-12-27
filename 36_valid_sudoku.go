package main

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	grid := [9][9]bool{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val, err := strconv.Atoi(string(board[row][col]))
			if err != nil {
				continue
			}
			val--
			gridIndex := col/3 + (row/3) * 3
			if rows[row][val] || cols[col][val] || grid[gridIndex][val] {
				return false
			}
			rows[row][val] = true
			cols[col][val] = true
			grid[gridIndex][val] = true
		}
	}
	return true
}