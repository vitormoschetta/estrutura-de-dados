package array

type Sudoku [9][9]int

var sudoku Sudoku

func Exec3() {
	sudoku[0][0] = 1
	sudoku[0][1] = 2
	sudoku[0][2] = 3
	sudoku[0][3] = 4
	sudoku[0][4] = 5
	sudoku[0][5] = 6
	sudoku[0][6] = 7
	sudoku[0][7] = 8
	sudoku[0][8] = 9

	sudoku[1][0] = 4
	sudoku[1][1] = 5
	sudoku[1][2] = 6
	sudoku[1][3] = 7
	sudoku[1][4] = 8
	sudoku[1][5] = 9
	sudoku[1][6] = 1
	sudoku[1][7] = 2
	sudoku[1][8] = 3

	sudoku[2][0] = 7
	sudoku[2][1] = 8
	sudoku[2][2] = 9
	sudoku[2][3] = 1
	sudoku[2][4] = 2
	sudoku[2][5] = 3
	sudoku[2][6] = 4
	sudoku[2][7] = 5
	sudoku[2][8] = 6

	sudoku[3][0] = 2
	sudoku[3][1] = 3
	sudoku[3][2] = 4
	sudoku[3][3] = 5
	sudoku[3][4] = 6
	sudoku[3][5] = 7
	sudoku[3][6] = 8
	sudoku[3][7] = 9
	sudoku[3][8] = 1

	sudoku[4][0] = 5
	sudoku[4][1] = 6
	sudoku[4][2] = 7
	sudoku[4][3] = 8
	sudoku[4][4] = 9
	sudoku[4][5] = 1
	sudoku[4][6] = 2
	sudoku[4][7] = 3
	sudoku[4][8] = 4

	sudoku[5][0] = 8
	sudoku[5][1] = 9
	sudoku[5][2] = 1
	sudoku[5][3] = 2
	sudoku[5][4] = 3
	sudoku[5][5] = 4
	sudoku[5][6] = 5
	sudoku[5][7] = 6
	sudoku[5][8] = 7

	sudoku[6][0] = 3
	sudoku[6][1] = 4
	sudoku[6][2] = 5
	sudoku[6][3] = 6
	sudoku[6][4] = 7
	sudoku[6][5] = 8
	sudoku[6][6] = 9
	sudoku[6][7] = 1
	sudoku[6][8] = 2

	sudoku[7][0] = 6
	sudoku[7][1] = 7
	sudoku[7][2] = 8
	sudoku[7][3] = 9
	sudoku[7][4] = 1
	sudoku[7][5] = 2
	sudoku[7][6] = 3
	sudoku[7][7] = 4
	sudoku[7][8] = 5

	sudoku[8][0] = 9
	sudoku[8][1] = 1
	sudoku[8][2] = 2
	sudoku[8][3] = 3
	sudoku[8][4] = 4
	sudoku[8][5] = 5
	sudoku[8][6] = 6
	sudoku[8][7] = 7
	sudoku[8][8] = 8

	PrintSudokuIn2D()
}

func PrintSudokuIn2D() {
	for _, row := range sudoku {
		for _, col := range row {
			print(col, " ")
		}
		println()
	}
}
