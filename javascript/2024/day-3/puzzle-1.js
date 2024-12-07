const fs = require("fs");

const INPUT_FILE = "./2024/day-3/input.txt";
const TEST_FILE = "./2024/day-3/test.txt";

const content = fs.readFileSync(INPUT_FILE).toString();
console.log(content);
const exp = /mul\((\d*),(\d*)\)/g;

const mults = content.matchAll(exp);

console.log(mults);

let result = 0;
for (let m of mults) {
  const n1 = Number(m[1]);
  const n2 = Number(m[2]);

  console.log(n1, "*", n2);
  result += n1 * n2;
}

console.log("Result =", result);
