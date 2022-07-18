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


  return (
    <div className="App">
      <header className="App-header">
        <h2>Sudoku Solver</h2>
        <div className="buttonContainer">
          <button className = "solveButton">Solve</button>
          <button className = "verifyButton">Verify</button>
          <button onClick = {() =>setSudokuArr(getDeepCopy(initial))}className = "clearButton">Clear</button>
        </div>
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

