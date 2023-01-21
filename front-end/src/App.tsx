import React, { useState } from 'react';
import './App.css';
import { initial, solveSudokuURL, verifySudokuURL} from "./constants";
import {getDeepCopy} from "./utils";
import { makepuzzle } from "sudoku";
import { SudokuBoard } from "./Components/SudokuBoard";
function App() {
  const [sudokuArr, setSudokuArr] = useState(getDeepCopy(initial));
  const [sudokuArrStart, setSudokuArrStart] = useState(getDeepCopy(initial));
  const [solvedSudokusArr, setSolvedSudokusArr] = useState(getDeepCopy(initial));
  const [unsolvedSudokuArr, setUnsolvedSudokusArr] = useState(null);

  function getNewSudoku() {
    const puzzle = makepuzzle();
  let grid = [];
  let row = [];
  for (let i = 0; i < puzzle.length; i++) {
    if (puzzle[i] === null) {
      puzzle[i] = -1;
    } else {
      puzzle[i]++;
    }
    row.push(puzzle[i]);
    if (row.length === 9) {
      grid.push(row);
      row = [];
    }
  }
    setSudokuArrStart(getDeepCopy(grid))
    setSudokuArr(getDeepCopy(grid))
  }

  function verifySudoku(input:[[number]]) {
    const xhttp = new XMLHttpRequest();
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
            setUnsolvedSudokusArr(getDeepCopy(sudokuArr))
            setSolvedSudokusArr(getDeepCopy(obj.Solutions[0]));
          } else {
            alert("No solutions for this sudoku")
          }
        }
        else{ 
          console.log("ERROR: " + xhttp.response)
        }
      }
    }
    xhttp.open("POST",solveSudokuURL);
    xhttp.send(JSON.stringify({grid:sudokuToSolve}));
  }

  const clearSolveGrid = () => {
    setSolvedSudokusArr(getDeepCopy(initial));
    setUnsolvedSudokusArr(null);
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
        {SudokuBoard(sudokuArr,setSudokuArr,sudokuArrStart)}
        <h4>Solve Sudoku:</h4>
        <div className="buttonContainer">
          <button onClick = {() => solveSudoku(solvedSudokusArr)} className = "solveButton">Solve</button>
          <button onClick = {() => verifySudoku(solvedSudokusArr)} className = "verifyButton">Verify</button>
          <button onClick = {() => clearSolveGrid()} className = "clearButton">Clear</button>
        </div>
        {SudokuBoard(solvedSudokusArr,setSolvedSudokusArr,unsolvedSudokuArr)}
        <div><p>Source Code on Github: <a href="https://github.com/Oliver-Looney/sudoku-solver">https://github.com/Oliver-Looney/sudoku-solver</a></p></div>
         {/*<div><p>My Portfolio: <a href="http://oliverlooney.com/">http://oliverlooney.com/</a></p></div>*/}
      </header>
    </div>
  );
}
export default App;