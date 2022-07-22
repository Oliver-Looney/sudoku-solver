import React, { useState } from 'react';
import './App.css';

let initial = [
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

function App() {
  const [sudokuArr, setSudokuArr] = useState(getDeepCopy(initial));

  function getDeepCopy(arr: number[][]) {
    return JSON.parse(JSON.stringify(arr));
  }

  function onInputChange(e: { target: { value: string; }; }, row:number, col:number){
    console.log("onInputChange")
    var val = parseInt(e.target.value) || -1, grid = getDeepCopy(sudokuArr);
    if (val === -1 || (val >= 1 && val <=9)){
      grid[row][col] = val;
    }
    setSudokuArr(grid);
  }

  function populateWithCompletedSudoku() {
    const grid = [
      [7, 8, 5, 6, 1, 2, 3, 9, 4],
      [9, 1, 4, 7, 8, 3, 2, 6, 5],
      [3, 6, 2, 4, 9, 5, 8, 1, 7],
      [6, 9, 1, 2, 7, 8, 5, 4, 3],
      [4, 3, 7, 1, 5, 6, 9, 2, 8],
      [2, 5, 8, 9, 3, 4, 1, 7, 6],
      [1, 2, 3, 5, 6, 7, 4, 8, 9],
      [8, 4, 6, 3, 2, 9, 7, 5, 1],
      [5, 7, 9, 8, 4, 1, 6, 3, 2]]
    setSudokuArr(getDeepCopy(grid))
  }

  function verifySudoku() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
      // document.getElementById("my-demo").innerHTML = this.responseText;
      // alert(this.responseText);
      if (xhttp.readyState === 4) {
        alert(xhttp.response);
        console.log(xhttp);
      }
    }
    // xhttp.open("POST","https://wki05pg2og.execute-api.eu-west-1.amazonaws.com/default/sudoku-solver-verify/sudoku-solver-verify");
    xhttp.open("POST","https://px21zcm9fi.execute-api.eu-west-1.amazonaws.com/default/sudoku-solver-verify");
    // xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhttp.send(JSON.stringify({grid:sudokuArr}));
    console.log(JSON.stringify({grid:sudokuArr}));
  }

  return (
    <div className="App">
      <header className="App-header">
        <h2>Sudoku Solver</h2>
        <div className="buttonContainer">
          <button className = "solveButton">Solve</button>
          <button onClick = {() => verifySudoku()} className = "verifyButton">Verify</button>
          <button onClick = {() =>setSudokuArr(getDeepCopy(initial))}className = "clearButton">Clear</button>
        </div>
        <button onClick = {() => populateWithCompletedSudoku() }>Populate sudoku</button>
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
            <td>

            </td>
          </tbody>
        </table>
      </header>
    </div>
  );
}

export default App;

function clearButton() {
  throw new Error('Function not implemented.');
}
// function initial(initial: any): [any, any] {
//   throw new Error('Function not implemented.');
// }

