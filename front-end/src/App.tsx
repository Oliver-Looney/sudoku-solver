import React, { useState } from 'react';
import './App.css';
import { initial, solveSudokuURL, verifySudokuURL} from "./constants";
import {getCellType, getDeepCopy} from "./utils";

function App() {
  const [sudokuArr, setSudokuArr] = useState(getDeepCopy(initial));
  const [solvedSudokusArr, setSolvedSudokusArr] = useState([getDeepCopy(initial)]);
  const [unsolvedSudokuArr, setUnolvedSudokusArr] = useState(null);
  function onInputChange(e: { target: { value: string; }; }, row:number, col:number){
    const val = parseInt(e.target.value) || -1, grid = getDeepCopy(sudokuArr);
    if (val === -1 || (val >= 1 && val <=9)){
      grid[row][col] = val;
    }
    setSudokuArr(grid);
  }

  function populateWithCompletedSudoku() {
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

  function solveSudoku() {
    const xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
      if (xhttp.readyState === 4) {
        if (xhttp.status === 200){
          let obj = JSON.parse(xhttp.response);
          if (obj.Solvable === true) {
            setUnolvedSudokusArr(getDeepCopy(sudokuArr))
            setSolvedSudokusArr(getDeepCopy(obj.Solutions));
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
    xhttp.send(JSON.stringify({grid:sudokuArr}));
  }

  return (
    <div className="App">
      <header className="App-header">
        <h2>Sudoku Solver</h2>
        <h4>Play Sudoku:</h4>
        <div className="buttonContainer">
          <button onClick = {() => solveSudoku()} className = "solveButton">Solve</button>
          <button onClick = {() => verifySudoku(sudokuArr)} className = "verifyButton">Verify</button>
          <button onClick = {() =>setSudokuArr(getDeepCopy(initial))} className = "clearButton">Clear</button>
        </div>
        <button onClick = {() => populateWithCompletedSudoku() }>Populate with example sudoku</button>
        <table>
          <tbody>
            {
              [0,1,2,3,4,5,6,7,8].map((row,rindex) => {
                return <tr key = {rindex} className ={(row + 1) % 3 === 0 ? 'bBorder': ''}>
                  {[0,1,2,3,4,5,6,7,8].map((col,cindex) => { 
                    return <td key = {rindex + cindex} className ={(col + 1) % 3 === 0 ? 'rBorder': ''}>
                    <input onChange = { (e) => onInputChange(e, row, col )}
                      value = {sudokuArr[row][col] === -1 ? '': sudokuArr[row][col]}
                      className="cellInput"
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
        <button onClick = {() => verifySudoku(solvedSudokusArr[0])} className = "verifyButton">Verify</button>
        <button onClick = {() => setSolvedSudokusArr([getDeepCopy(initial)])} className = "clearButton">Clear</button>
        </div>
        <table>
          <tbody>
            {
              [0,1,2,3,4,5,6,7,8].map((row,rindex) => {
                return <tr key = {rindex} className ={(row + 1) % 3 === 0 ? 'bBorder': ''}>
                  {[0,1,2,3,4,5,6,7,8].map((col,cindex) => { 
                    return <td key = {rindex + cindex} className ={(col + 1) % 3 === 0 ? 'rBorder': ''}>
                    <input 
                      value = {solvedSudokusArr[0][row][col] === -1 ? '': solvedSudokusArr[0][row][col]}
                      className ={getCellType(solvedSudokusArr[0],unsolvedSudokuArr,row,col)}
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