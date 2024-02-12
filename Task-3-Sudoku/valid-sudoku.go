// APPROACH :- BackTracking for Solving the sudoku and Maps for verifying the solution
// restriction :- 1. no repeated digits across row
// 				  2. no repeated digits across column
// 				  3. no repeated digits across smaller 3x3 digits

// Solution :-
//  	step 1 :- using maps to map frequency of digits across row and column and 3x3 matrix if any element in
// 				  map has frequency > 1 the matrix is not valid and the element of that map can give us
// 				  the row and column for the wrong digit on the sudoku

//      step 2 :- if the matrix is valid than checking for if the sudoku is solvable or not
// 	    step 3 :- for this i first gather the available digit at the index and if multiple digit are present
// 				  than the recursive approach with backtracking came into place
//  	step 4 :- if at the end of loops if still we cant find any digit to fit properly we can say the matrix
// 				  is valid but not solvable else we can say the matrix is solvable and valid




package main

import "fmt"


// checking for frequency of digit across sudoku row for checking validitiy of the sudoku
// @params :- a single row of the sudoku
// @returns :- a bool giving if sudoku is valid or not and a int regarding on which row or column the element fail
func isRowValid(row []int) (bool, int) {

	m := make(map[int]int)
	// loop through the map and calculate the occurance of each digit if > 1 the sudoku is not valid

	for i, v := range row {
		if v != 0 {
			m[v]++
			if m[v] > 1 {
				return false, i
			}
		}
	}

	return true, -1
}


// transpose swap the element of row and column so i can use funcions for row as funcitons for column
func transpose(board [][]int) {
	
	var temp int

	for i := 0; i < len(board); i++ {
		for j := i; j < len(board[i]); j++ {
			temp = board[i][j]
			board[i][j] = board[j][i]
			board[j][i] = temp
		}
	}

}


// check for validity of small 3x3 matrix returns true if valid else return false and index at which point the
//  sudoku fails
// @params :- sudoku it self and index of first element of the sub 3x3 matrix
// @returns :- returns a bool validity of 3x3 matrix and index of the wrong digit
func isValidMatrix(matrix [][]int, row int, column int) (bool, int, int) {
	
	m := make(map[int]int)

	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			
			if matrix[i][j] != 0 {
				m[matrix[i][j]]++
				if m[matrix[i][j]] > 1 {
					return false, i, j
				}
			}

		}
	}

	return true, -1, -1
}


// check for validity of sudoku is the given sudoku valid or not
// @params :- a sudoku it self for checking for validity
// @returns :- return boolean regarding validity of matrix else return the index where the element fails
func isValidSudoku(board [][]int) (bool, int, int) {
	var invalidRow int
	var invalidColumn int
	for i, v := range board {
		if valid, index := isRowValid(v); !valid {
			invalidRow = i
			invalidColumn = index
			return false, invalidRow, invalidColumn
		}
	}
	transpose(board)
	for i, v := range board {
		if valid, index := isRowValid(v); !valid {
			invalidColumn = i
			invalidRow = index
			return false, invalidRow, invalidColumn
		}
	}
	transpose(board)
	var counterRow, counterColumn int = 0, 0
	var firstIndex, lastIndex int = 0, 0
	for counter := 0; counter < 9; counter++ {
		firstIndex = counterRow * 3
		lastIndex = counterColumn * 3
		if valid, row, column := isValidMatrix(board, lastIndex, firstIndex); !valid {
			return false, column, row
		}
		counterRow++
		counterRow = counterRow % 3
		if counterRow == 0 {
			counterColumn++
		}

	}
	return true, -1, -1
}

// check occurance of each digit across row and column of sudoku for solvability purpose
// @params :- a array of sudoku digits either across row or either across column
// @returns :- a map with occurance of digit
func rowMap(row []int) map[int]int {
	m := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}

	for _, v := range row {
		if _, ok := m[v]; ok {
			m[v]++
		}
	}

	return m
}


// function for calculating the frequency of each digit on a 3x3 matrix for solvability checking
// @params :- sudoku it self and starting index of the first element of sub 3x3 sudoku
// @returns :- a map for frequency of occured digit across a valid 3x3 matrix
func matrixMap(sudoku [][]int, row int, column int) map[int]int {

	m := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}

	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if _, ok := m[sudoku[i][j]]; ok {
				m[sudoku[i][j]]++
			}
		}
	}

	return m
}

// debug ## a debug function else for printing the sudoku
func print(sudoku [][]int){
	for _,row := range sudoku {
		for _,column := range row {
			fmt.Print(column, " ")
		}
		fmt.Println()
	}
}

// actually solves sudoku using backtracking untill it find sudo is not solvable or solves the sudoku
// @params :- a sudoku it self with containing element or 0
// @returns :- a bool regarding either the sudoku is solvable or not
func solveSudoku(sudoku [][]int) bool {
	// variable for showing actually available input for putting in place in sudoku
	mapRow := make(map[int]int)
	mapColumn := make(map[int]int)
	mapMatrix := make(map[int]int)
	var availableAnswer []int // store possible solution at a point in sudoku

	for i, row := range sudoku {
		for j, _ := range row {

			// finding first occurance of non zero element
			if sudoku[i][j] != 0{
				continue
			}

			// calculating solution map that shows possible input across row column and matrix
			mapRow = rowMap(row)
			transpose(sudoku)
			mapColumn = rowMap(sudoku[j])
			transpose(sudoku)
			mapMatrix = matrixMap(sudoku, (i/3)*3, (j/3)*3)


			// storing available solution in a array
			for available := 1; available <= 9; available++ {
				// if the element is not occurade in row,column and matrix the digit is available to use
				if mapRow[available]+mapColumn[available]+mapMatrix[available] == 0 {
					availableAnswer = append(availableAnswer, available)
				}
			}

			if len(availableAnswer) == 0 {
				// condition for stopping backtrack usefull at last index of sudoku
				return false
			} else {

				var index = 0   //for looping through possible answer array

				for index < len(availableAnswer) {
					// set the available answer in place of i,j
					sudoku[i][j]=availableAnswer[index]

					// continuing the solution with one of the value from multiple possible answer 
					// recursion
					var isSolvable = solveSudoku(sudoku)
					if isSolvable {
						return true
					}else{
						// if the answer not solve sudoku we reset the element and try another one
						// used in backtrack
						sudoku[i][j]=0
						index++
					}
				}
				// if any possible answer dont satisfy sudoku we simple backtrac
				return false
			}
		}
	}
	// if all digits are set we simply solved the sudoku so completing the backtrack
	return true
}

func main() {

    // [5, 3, 4, 6, 7, 8, 9, 1, 2],
    // [6, 7, 2, 1, 9, 5, 3, 4, 8],
    // [1, 9, 8, 3, 4, 2, 5, 6, 7],
    // [8, 5, 9, 7, 6, 1, 4, 2, 3],
    // [4, 2, 6, 8, 5, 3, 7, 9, 1],
    // [7, 1, 3, 9, 2, 4, 8, 5, 6],
    // [9, 6, 1, 5, 3, 7, 2, 8, 4],
    // [2, 8, 7, 4, 1, 9, 6, 3, 5],
    // [3, 4, 5, 2, 8, 6, 1, 7, 9]

	sudoku := [][]int{
		{0, 3, 4, 0, 7, 8, 0, 0, 0},
		{6, 7, 2, 1, 9, 0, 0, 0, 0},
		{1, 0, 8, 3, 4, 2, 0, 0, 0},
		{0, 5, 9, 0, 6, 1, 0, 0, 0},
		{0, 2, 6, 8, 5, 3, 0, 0, 0},
		{7, 1, 3, 0, 2, 0, 0, 0, 0},
		{9, 6, 0, 5, 3, 7, 0, 8, 0},
		{2, 0, 7, 4, 1, 9, 0, 0, 0},
		{3, 4, 5, 0, 8, 6, 0, 0, 0},
	}

	// sudoku := [][]int{
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// 	{0,0, 0, 0, 0, 0,0, 0,0},
	// }

	// sudoku := [][]int{
    //     {5, 3, 4, 6, 7, 8, 0, 1, 2},
    //     {6, 7, 2, 1, 0, 5, 9, 4, 8},
    //     {1, 9, 8, 3, 4, 2, 5, 6, 7},
    //     {8, 5, 9, 7, 6, 1, 4, 2, 3},
    //     {4, 2, 6, 8, 5, 3, 7, 9, 1},
    //     {7, 1, 3, 9, 2, 4, 8, 5, 6},
    //     {9, 6, 1, 5, 3, 7, 2, 8, 4},
    //     {2, 8, 7, 4, 1, 9, 6, 3, 5},
    //     {3, 4, 5, 2, 8, 6, 1, 7, 9},
    // }

	// checking for validity of sudoku or the index where the sudoku fails
	valid, row, column := isValidSudoku(sudoku)

	if !valid {
		fmt.Println("invalid row is ", row, " and invalid column is ", column)
	} else {
		// if sudoku is valid than trying to solve the sudoku
		if solveSudoku(sudoku) {
			if valid2,_,_ := isValidSudoku(sudoku); valid2 {
				fmt.Println("sudoku is solvable and valid")
			}
			print(sudoku)
		} else {
			fmt.Println("the sudoku can't be solvable")
		}

	}
}
