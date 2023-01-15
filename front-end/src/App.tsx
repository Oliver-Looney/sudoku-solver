import React, { useState } from 'react';
import './App.css';
import { initial, solveSudokuURL, verifySudokuURL} from "./constants";
import {getCellType, getDeepCopy} from "./utils";

function App() {
  const [sudokuArr, setSudokuArr] = useState(getDeepCopy(initial));
  const [sudokuArrStart, setSudokuArrStart] = useState(getDeepCopy(initial));
  const [solvedSudokusArr, setSolvedSudokusArr] = useState(getDeepCopy(initial));
  const [unsolvedSudokuArr, setUnolvedSudokusArr] = useState(null);

  function onInputChangePlay(e: { target: { value: string; }; }, row:number, col:number, setGrid: (arg0: number[][]) => void, currentGrid: number[][], originalGrid: number[][]|null){
    const val = parseInt(e.target.value) || -1, grid = getDeepCopy(currentGrid);
    // if (originalGrid === null){
      // grid[row][col] = val;
    // } else
    if ((val === -1 || (val >= 1 && val <=9)) && (originalGrid === null || originalGrid[row][col] === -1)){
      grid[row][col] = val;
    }
    setGrid(grid);
  }

  function getNewSudoku() {
    const grid = [
      [7, -1, 5, 6, 1, 2, 3, 9, 4],
      [-1, -1, 4, 7, -1, 3, 2, 6, 5],
      [3, -1, -1, 4, -1, -1, -1, -1, -1],
      [-1, -1, 1, -1, 7, 8, 5, -1, -1],
      [-1, -1, 7, -1, 5, 6, 9, 2, 8],
      [2, 5, -1, -1, 3, -1, 1, -1, -1],
      [-1, 2, 3, -1, -1, 7, 4, 8, -1],
      [-1, -1, 6, 3, 2, 9, 7, -1, 1],
      [5, -1, 9, -1, 4, -1, -1, -1, 2]]
    setSudokuArrStart(getDeepCopy(grid))
    setSudokuArr(getDeepCopy(grid))
  }

  function verifySudoku(input:[[number]]) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
      if (xhttp.readyState === 4) {
        alert(xhttp.response);
      }
    }
    xhttp.open("POST",verifySudokuURL);
    xhttp.send(JSON.stringify({grid:input}));
  }

  function solveSudoku(sudokuToSolve: number[][]) {
    const xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
      if (xhttp.readyState === 4) {
        if (xhttp.status === 200){
          let obj = JSON.parse(xhttp.response);
          if (obj.Solvable === true) {
            setUnolvedSudokusArr(getDeepCopy(sudokuArr))
            setSolvedSudokusArr(getDeepCopy(obj.Solutions[0]));
          } else {
            alert("No solutions for this sudoku")
          }
        }
        else{ 
          alert("ERROR: " + xhttp.response)
        }
      }
    }
    xhttp.open("POST",solveSudokuURL);
    xhttp.send(JSON.stringify({grid:sudokuToSolve}));
  }

  const clearSolveGrid = () => {
    setSolvedSudokusArr(getDeepCopy(initial));
    setUnolvedSudokusArr(null);
  }

  return (
    <div className="App">
      <header className="App-header">
        <h2>Sudoku Solver</h2>
        <h4>Play Sudoku:</h4>
        <div className="buttonContainer">
          <button onClick = {() => solveSudoku(sudokuArrStart)} className = "solveButton">Solve</button>
          <button onClick = {() => verifySudoku(sudokuArr)} className = "verifyButton">Verify</button>
          <button onClick = {() => setSudokuArr(getDeepCopy(sudokuArrStart))} className = "clearButton">Clear</button>
        </div>
        <button onClick = {() => getNewSudoku() } className="playButton">Play A New Sudoku</button>
        <table>
          <tbody>
            {
              [0,1,2,3,4,5,6,7,8].map((row,rindex) => {
                return <tr key = {rindex} className ={(row + 1) % 3 === 0 ? 'bBorder': ''}>
                  {[0,1,2,3,4,5,6,7,8].map((col,cindex) => { 
                    return <td key = {rindex + cindex} className ={(col + 1) % 3 === 0 ? 'rBorder': ''}>
                    <input onChange = { (e) => onInputChangePlay(e, row, col, setSudokuArr, sudokuArr, sudokuArrStart )}
                      value = {sudokuArr[row][col] === -1 ? '': sudokuArr[row][col]}
                        className ={getCellType(sudokuArr,sudokuArrStart,row,col)}
                      />
                  </td>
                  })}
                </tr>
              })
            }
          </tbody>
        </table>
        <h4>Solve Sudoku:</h4>
        <div className="buttonContainer">
          <button onClick = {() => solveSudoku(solvedSudokusArr)} className = "solveButton">Solve</button>
          <button onClick = {() => verifySudoku(solvedSudokusArr)} className = "verifyButton">Verify</button>
          <button onClick = {() => clearSolveGrid()} className = "clearButton">Clear</button>
        </div>
        <table>
          <tbody>
            {
              [0,1,2,3,4,5,6,7,8].map((row,rindex) => {
                return <tr key = {rindex} className ={(row + 1) % 3 === 0 ? 'bBorder': ''}>
                  {[0,1,2,3,4,5,6,7,8].map((col,cindex) => { 
                    return <td key = {rindex + cindex} className ={(col + 1) % 3 === 0 ? 'rBorder': ''}>
                    <input onChange = { (e) => onInputChangePlay(e, row, col, setSolvedSudokusArr, solvedSudokusArr, unsolvedSudokuArr )}
                      value = {solvedSudokusArr[row][col] === -1 ? '': solvedSudokusArr[row][col]}
                      className ={getCellType(solvedSudokusArr,unsolvedSudokuArr,row,col)}
                      />
                  </td>
                  })}
                </tr>
              })
            }
          </tbody>
        </table>
        <div><p>Source Code on Github: <a href="https://github.com/Oliver-Looney/sudoku-solver">https://github.com/Oliver-Looney/sudoku-solver</a></p></div>
         {/*<div><p>My Portfolio: <a href="http://oliverlooney.com/">http://oliverlooney.com/</a></p></div>*/}
      </header>
    </div>
  );
}

export default App;