const fs = require("fs");

const INPUT_FILE = "./2024/day-4/input.txt";
const TEST_FILE = "./2024/day-4/test.txt";

const strToArr2D = (str) => str.split("\n").map((s) => s.split(""));

const content = fs.readFileSync(INPUT_FILE).toString();

const findWord = (str, word) => {
  let findings = 0;
  for (let i = 0; i < str.length; i++) {
    if (str[i] === word[0]) {
      let wordIndex = 0;
      for (let k = i; k < str.length; k++) {
        if (str[k] === word[wordIndex]) {
          wordIndex++;

          // We found the last letter
          if (wordIndex === word.length) {
            findings++;
            break;
          }
        }
      }
    }
  }

  return findings;
};

const countWords = (arr2D, word) => {
  let findings = 0;
  for (let i = 0; i < arr2D.length; i++) {
    for (let j = 0; j < arr2D[i].length; j++) {
      if (arr2D[i][j] === word[0]) {
        // Find HORIZONTAL
        let wordIndex = 0;
        for (let k = j; k < arr2D[i].length; k++) {
          if (arr2D[i][k] === word[wordIndex]) {
            wordIndex++;

            // We found the last letter
            if (wordIndex === word.length) {
              findings++;
              break;
            }
          }
        }

        // Find VERTICAL
        wordIndex = 0;
        for (let k = i; k < arr2D.length; k++) {
          if (arr2D[k][j] === word[wordIndex]) {
            wordIndex++;

            // We found the last letter
            if (wordIndex === word.length) {
              findings++;
              break;
            }
          }
        }
      }
    }
  }
  return findings;
};

const contentWords = countWords(strToArr2D(content), "XMAS");
console.log("Found:", contentWords);

let findings = content.split("\n").map((l) => ({
  line: l,
  findings: findWord(l, "XMAS"),
  word: "XMAS",
}));

findings = findings.concat(
  content.split("\n").map((l) => ({
    line: l,
    findings: findWord(l, "SAMX"),
    word: "SAMX",
  }))
);

console.log("Findings", findings);

const countWordsInArr2D = (arr2D, word) => {
  let findings = 0;

  for (let row = 0; row < arr2D.length; row++) {
    for (let col = 0; col < arr2D[row].length; col++) {
      // If the first letter is found, we look for the word
      // horizontaly, diagonally and vertically
      if (arr2D[row][col] === word[0]) {
        // HORIZONTAL SEARCH. Move across the Columns
        let wordIndex = 0;
        for (let k = col; k < arr2D[row].length; k++) {
          if (arr2D[row][k] !== word[wordIndex]) {
            break;
          }

          wordIndex++;

          // Whole word is found
          if (wordIndex >= word.length) {
            console.log(
              "Found horizontal ",
              word,
              "from (",
              row,
              ",",
              col,
              ")"
            );
            findings++;
            break;
          }
        }

        // VERTICAL SEARCH. Move across the ROWS
        wordIndex = 0;
        for (let k = row; k < arr2D.length; k++) {
          if (arr2D[k][col] !== word[wordIndex]) {
            break;
          }
          wordIndex++;

          // Whole word is found
          if (wordIndex >= word.length) {
            console.log("Found vertical ", word, "from (", row, ",", col, ")");

            findings++;
            break;
          }
        }

        // DIAGONAL SEARCH. Move across the ROWS and COLUMNS
        wordIndex = 0;
        for (
          let k = 0;
          k + row < arr2D.length && k + col < arr2D[row].length;
          k++
        ) {
          if (arr2D[row + k][col + k] !== word[wordIndex]) {
            break;
          }
          wordIndex++;

          // Whole word is found
          if (wordIndex >= word.length) {
            console.log("Found Diagonal ", word, "from (", row, ",", col, ")");

            findings++;
            break;
          }
        }

        wordIndex = 0;
        for (let k = 0; k + row < arr2D.length && col - k >= 0; k++) {
          if (arr2D[row + k][col - k] !== word[wordIndex]) {
            break;
          }
          wordIndex++;

          // Whole word is found
          if (wordIndex >= word.length) {
            console.log(
              "Found Reverse Diagonal ",
              word,
              "from (",
              row,
              ",",
              col,
              ")"
            );

            findings++;
            break;
          }
        }
      }
    }
  }

  return findings;
};

console.log(
  "countWordsInArr2D:",
  countWordsInArr2D(strToArr2D(content), "XMAS") +
    countWordsInArr2D(strToArr2D(content), "SAMX")
);
