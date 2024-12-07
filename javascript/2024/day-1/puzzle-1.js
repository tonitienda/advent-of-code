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

console.log("----");
l1.sort();
l2.sort();

console.log(l1);
console.log(l2);

const result = l1
  .map((n1, idx) => Math.abs(l2[idx] - n1))
  .reduce((n, acc) => acc + n, 0);

console.log("Result:", result);
