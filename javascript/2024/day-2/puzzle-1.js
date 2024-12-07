const fs = require("fs");
const { getNumber2DArray } = require("../../utils/file");

const negative = (n) => n < 0;
const positive = (n) => n > 0;
const withinRange = (lower, upper) => (n) => n <= upper && n >= lower;

const isSafe = (line) =>
  line.every(withinRange(-3, 3)) &&
  (line.every(negative) || line.every(positive)) &&
  line.every((n) => n != 0);

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
  .map((line) =>
    line.reduce((steps, number, idx) => {
      if (idx >= line.length - 1) {
        return steps;
      }

      steps[idx] = number - line[idx + 1];
      return steps;
    }, [])
  )
  .map(isSafe)
  .reduce(makeCounter, {});

console.log("Result =", result);
