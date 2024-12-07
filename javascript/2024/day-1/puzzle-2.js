const fs = require("fs");

const input = fs.readFileSync("./2024/day-1/input.txt", "utf8");
const test = fs.readFileSync("./2024/day-1/test.txt", "utf8");

const SEPARATOR = "   ";

const contents = input;

const [l1, l2] = contents
  .split("\n")
  .map((line) => line.split(SEPARATOR))
  .reduce(
    ([l1, l2], [p1, p2]) => [
      [...l1, Number(p1)],
      [...l2, Number(p2)],
    ],
    [[], []]
  );
console.log(l1);
console.log(l2);

const l2map = l2.reduce((acc, n) => {
  if (acc[n]) {
    acc[n]++;
  } else {
    acc[n] = 1;
  }
  return acc;
}, {});

console.log("----");
console.log("PUZZLE 2");
console.log(l2map);

// console.log(l1);
// console.log(l2);

const similarityScores = l1.map((n) => (l2map[n] || 0) * n);

console.log("similarityScores", similarityScores);

const result = similarityScores.reduce((n, acc) => acc + n, 0);

console.log("Result:", result);
