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

const getDeepCopy = (arr: number[][]) => JSON.parse(JSON.stringify(arr));

const APIGatewayURL = "https://px21zcm9fi.execute-api.eu-west-1.amazonaws.com/default/"

const verifySudokuURL = APIGatewayURL + "sudoku-solver-verify"
const solveSudokuURL = APIGatewayURL + "sudoku-solver-solve"

export { initial, getDeepCopy, verifySudokuURL, solveSudokuURL }