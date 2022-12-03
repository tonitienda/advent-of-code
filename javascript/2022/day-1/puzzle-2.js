const file = require("../../utils/file");
const { sum } = require("../../utils/array-extensions");

const path = require("path");

const data = file
  .getLines(path.join(__dirname, "input.txt"), "\n\n")
  .map((foods) => sum(foods.split("\n").map(Number)))
  .sort((a, b) => (a > b ? -1 : 1));

console.log(data[0] + data[1] + data[2]);
