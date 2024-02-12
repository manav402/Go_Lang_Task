**PROBLEM** 
  1. check if the given sudoku is in a valid format or not
  2. if the sudoku is invalid return the index of the element where the sudoku fails
  3. if the sudoku is valid see if the solution is possible or not

**RESTRICTION**
  1. no repeated digits across the row
  2. no repeated digits across the column
  3. no repeated digits across smaller 3x3 digits

**APPROACH**
  - Backtracking for Solving the sudoku and Maps for verifying the solution

**STEPS**
   1. using maps to map the frequency of digits across rows and column and 3x3 matrix if any element in
 				  the map has frequency > 1 the matrix is not valid and the element of that map can give us
 				  the row and column for the wrong digit on the sudoku
   2. if the matrix is valid then check for if the sudoku is solvable or not
   3. for this I first gather the available digits at the index and if multiple digits are present
				  then the recursive approach with backtracking came into place
   4. if at the end of loops still we can't find any digit to fit properly we can say the matrix
				  is valid but not solvable else we can say the matrix is solvable and valid
  
