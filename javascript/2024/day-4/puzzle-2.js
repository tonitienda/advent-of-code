const fs = require("fs");

const INPUT_FILE = "./2024/day-4/input.txt";
const TEST_FILE = "./2024/day-4/test.txt";

const strToArr2D = (str) => str.split("\n").map((s) => s.split(""));

const content = fs.readFileSync(TEST_FILE).toString();

const Combinations = [
  [
    ["M", ".", "M"],
    [".", "A", "."],
    ["S", ".", "S"],
  ],
  [
    ["M", ".", "S"],
    [".", "A", "."],
    ["M", ".", "S"],
  ],
  [
    ["S", ".", "M"],
    [".", "A", "."],
    ["S", ".", "M"],
  ],
  [
    ["S", ".", "S"],
    [".", "A", "."],
    ["M", ".", "M"],
  ],
];

const findMatrixWithinMatrix = (arr2D, row, col, pattern, wildcard = ".") => {
  if (
    row + pattern.length > arr2D.length ||
    col + pattern[0].length > arr2D[0].length
  ) {
    return false;
  }

  for (let i = 0; i < pattern.length && i + row < arr2D.length; i++) {
    for (let j = 0; j < pattern[i].length && j + col < arr2D[i].length; j++) {
      if (
        pattern[i][j] !== wildcard &&
        pattern[i][j] !== arr2D[row + i][col + j]
      ) {
        return false;
      }
    }
  }

  return true;
};

console.log(content);

const arr = strToArr2D(content);

let count = 0;
for (let i = 0; i < arr.length; i++) {
  for (let j = 0; j < arr[i].length; j++) {
    for (let c = 0; c < Combinations.length; c++) {
      if (findMatrixWithinMatrix(arr, i, j, Combinations[c])) {
        console.log("Found pattern ", c, "in (", i, ",", j, ")");
        count++;
        break;
      }
    }
  }
}
console.log(count);

// let found = false;
// let count = 0;
// for (let row = 0; row < content.length - 2; row++) {
//   for (let col = 0; col < content[row].length - 2; col++) {
//     for (let i = 0; i < Combinations.length; i++) {
//       console.log();

//       found = false;
//       for (let row2 = 0; row2 < Combinations[i].length; row2++) {
//         for (let col2 = 0; col2 < Combinations[i][row2].length; col2++) {
//           let items = Combinations[i][row2].split("");
//           process.stdout.write(content[row + row2][col + col2]);
//           if (items[col2] === ".") {
//             continue;
//           }
//           if (
//             content[row + row2][col + col2] == items[col2] &&
//             row2 === 2 &&
//             col2 === 2
//           ) {
//             found = true;
//           }
//         }
//         console.log();
//       }
//       if (found) {
//         console.log("Found");
//         count++;
//         break;
//       }
//     }
//   }
// }

// console.log("Result", count);
