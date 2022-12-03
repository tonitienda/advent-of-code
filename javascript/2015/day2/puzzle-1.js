const file = require("../../utils/file");
const path = require("path");

const contents = file.getNumber2DArray(
  path.join(__dirname, "input.txt"),
  "\n",
  "x"
);

console.log(contents);

const boxSides = contents.map(([l, w, h]) => [l * w, l * h, w * h]);

const totalPaperPerBox = boxSides.map(
  (sides) =>
    sides.reduce((total, side) => total + 2 * side, 0) + Math.min(...sides)
);

const totalPaper = totalPaperPerBox.reduce(
  (total, boxPaper) => total + boxPaper
);

console.log(boxSides);
console.log(totalPaperPerBox);
console.log(totalPaper);
