const fs = require("fs");
const { getNumber2DArray } = require("../../utils/file");

const withinRange = (lower, upper) => (n) => n <= upper && n >= lower;

const lineWithout = (line, idx) => [
  ...line.slice(0, idx),
  ...line.slice(idx + 1),
];

const isSafeWithDampener = (line) => {
  if (isSafe(line)) {
    return true;
  }

  for (let idx = 0; idx < line.length; idx++) {
    const line2 = lineWithout(line, idx);

    console.log(line);
    console.log(line2, idx);

    if (isSafe(line2)) {
      return true;
    }
  }

  return false;
};
const isSafe = (line) => {
  let sign = 0;

  for (let idx = 0; idx < line.length - 1; idx++) {
    const current = line[idx];
    const next = line[idx + 1];

    const diff = current - next;

    if (sign == 0) {
      sign = diff;
    }

    if (diff == 0 || !withinRange(-3, 3)(diff) || diff * sign < 0) {
      return false;
    }
  }
  return true;
};

const makeCounter = (counter, value) => {
  if (!counter[value]) {
    counter[value] = 0;
  }

  counter[value]++;

  return counter;
};

const INPUT_FILE = "./2024/day-2/input.txt";
const TEST_FILE = "./2024/day-2/test.txt";

const ROW_SEPARATOR = "\n";
const CELL_SEPARATOR = " ";

const numbers = getNumber2DArray(INPUT_FILE, ROW_SEPARATOR, CELL_SEPARATOR);

const result = numbers
  .map((line) => isSafeWithDampener(line))
  .reduce(makeCounter, {});

console.log(result);
