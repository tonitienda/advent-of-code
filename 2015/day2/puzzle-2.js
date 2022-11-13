const file = require("../../utils/file");
const arr = require("../../utils/array-extensions");
const path = require("path");

const contents = file.getNumber2DArray(
  path.join(__dirname, "input.txt"),
  "\n",
  "x"
);

console.log(contents);

const ribbonPerBox = contents
  .map((measures) => arr.sort(measures))
  .map((measures) => measures[0] * 2 + measures[1] * 2 + arr.mult(measures));
console.log(ribbonPerBox);

const totalRibbon = arr.sum(ribbonPerBox);
console.log(totalRibbon);
