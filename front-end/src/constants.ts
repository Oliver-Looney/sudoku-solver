const initial = [
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1],
    [-1,-1,-1,-1,-1,-1,-1,-1,-1]
]

const APIGatewayURL = "https://rz1uamskd6.execute-api.eu-west-1.amazonaws.com/default/"

const verifySudokuURL = APIGatewayURL + "sudoku-solver-verify"
const solveSudokuURL = APIGatewayURL + "sudoku-solver-solve"

export { initial, verifySudokuURL, solveSudokuURL }