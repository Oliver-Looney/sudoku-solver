import React from 'react';
import {getCellType, getDeepCopy} from "../utils";

function onInputChangePlay(e: { target: { value: string; }; }, row:number, col:number, setGrid: (arg0: number[][]) => void, currentGrid: number[][], originalGrid: number[][]|null){
    const val = isNaN(parseInt(e.target.value)) ? -1 : parseInt(e.target.value);
    if ((val === -1 || (val >= 1 && val <=9)) && (originalGrid === null || originalGrid[row][col] === -1)){
        const grid = getDeepCopy(currentGrid);
        grid[row][col] = val;
        setGrid(grid)
    }

}
const SudokuBoard = (sudoku:number[][], setSudoku: (arg0: number[][]) => void, originalSudoku: number[][]|null) => {
    return(<div>
        <table>
            <tbody>
            {
                [0,1,2,3,4,5,6,7,8].map((row,rindex) => {
                    return <tr key = {rindex} className ={(row + 1) % 3 === 0 ? 'bBorder': ''}>
                        {[0,1,2,3,4,5,6,7,8].map((col,cindex) => {
                            return <td key = {rindex + cindex} className ={(col + 1) % 3 === 0 ? 'rBorder': ''}>
                                <input type="number" onChange = { (event) => onInputChangePlay(event, row, col, setSudoku, sudoku, originalSudoku )}
                                       value = {sudoku[row][col] === -1 ? '': sudoku[row][col]}
                                       className ={getCellType(sudoku,originalSudoku,row,col)}
                                />
                            </td>
                        })}
                    </tr>
                })
            }
            </tbody>
        </table>
    </div>)
}

export { SudokuBoard }