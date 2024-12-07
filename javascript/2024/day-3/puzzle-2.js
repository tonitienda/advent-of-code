const fs = require("fs");

const INPUT_FILE = "./2024/day-3/input.txt";
const TEST_FILE = "./2024/day-3/test.txt";

const content = fs.readFileSync(INPUT_FILE).toString();
console.log(content);

const exp = /mul\((\d*),(\d*)\)/g;

const multiply = (op) => {
  const mults = op.matchAll(exp);

  let result = 0;
  for (let m of mults) {
    const n1 = Number(m[1]);
    const n2 = Number(m[2]);

    result += n1 * n2;
  }

  return result;
};

// TODO - USe regex
const result = content
  .split("do()")
  .map((op) =>
    op.indexOf("don't()") > -1 ? op.substring(0, op.indexOf("don't()")) : op
  )
  .map(multiply)
  .reduce((acc, n) => acc + n, 0);

console.log(result);
