const fs = require("fs");

const INPUT_FILE = "./2024/day-7/input.txt";
const TEST_FILE = "./2024/day-7/test.txt";

const content = fs.readFileSync(INPUT_FILE).toString();

const Operations = [(n1, n2) => n1 + n2, (n1, n2) => n1 * n2];

const equations = content
  .split("\n")
  .map((l) => [
    Number(l.split(":")[0]),
    l.split(":")[1].trim().split(" ").map(Number),
  ]);

console.log(equations);

const findResult = (expected, current, numbers) => {
  if (!current) {
    current = numbers[0];
    numbers.shift();
  }

  if (numbers.length === 0 || current > expected) {
    return current;
  }

  const [next, ...rest] = numbers;

  for (let op of Operations) {
    const newCurrent = op(current, next);

    if (findResult(expected, newCurrent, rest) === expected) {
      return expected;
    }
  }

  return current;
};

const correctEqs = equations.filter((e) => findResult(e[0], 0, e[1]) === e[0]);
console.log(correctEqs);

const result = correctEqs.reduce((acc, item) => acc + item[0], 0);

console.log(result);
