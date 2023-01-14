const getDeepCopy = (arr: number[][]) => JSON.parse(JSON.stringify(arr));

const getCellType = (solvedSudoku: number[][], unsolvedSudoku:number[][]|null, row:number, col:number) => {
    if (unsolvedSudoku === null){
        return 'cellInput'
    }
    return solvedSudoku[row][col] !== unsolvedSudoku[row][col] || solvedSudoku[row][col] === -1 ? 'cellInput': 'userInputCell'
}

export {  getDeepCopy, getCellType }
