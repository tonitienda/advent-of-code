const file = require("../../utils/file");
const { sum } = require("../../utils/array-extensions");

const path = require("path");

function isRightOrder([left, right]) {
  //console.log(left, "<", right, "?");

  for (let i = 0; i < left.length; i++) {
    leftItem = left[i];

    if (i >= right.length) {
      console.log("RIGHT ORDER");
      return true;
    }

    rightItem = right[i];

    console.log(leftItem, "<=", rightItem, "?");

    if (Array.isArray(leftItem) && Array.isArray(rightItem)) {
      const result = isRightOrder([leftItem, rightItem]);
      if (!result) {
        console.log("NOT RIGHT ORDER");
        return false;
      }
      continue;
    }

    if (Array.isArray(leftItem) && !Array.isArray(rightItem)) {
      const result = isRightOrder([leftItem, [rightItem]]);
      if (!result) {
        console.log("NOT RIGHT ORDER");
        return false;
      }
      continue;
    }

    if (!Array.isArray(leftItem) && Array.isArray(rightItem)) {
      const result = isRightOrder([[leftItem], rightItem]);
      if (!result) {
        console.log("NOT RIGHT ORDER");
        return false;
      }
      continue;
    }

    if (!Array.isArray(leftItem) && !Array.isArray(rightItem)) {
      const result = leftItem <= rightItem;
      if (!result) {
        console.log("NOT RIGHT ORDER");
        return false;
      }
      continue;
    }
  }
  console.log("RIGHT ORDER");
  return true;
}

const data = file
  .getLines(path.join(__dirname, "test.txt"), "\n\n")
  .map((str) => str.split("\n"))
  .map(([left, right]) => [eval(left), eval(right)]);

const totalInRightOrder = data.reduce(
  (acc, row, idx) =>
    console.log(row) || (isRightOrder(row) ? acc + idx + 1 : acc),
  0
);

console.log(totalInRightOrder);

// console.log(data[0][0]);
// console.log(isRightOrder(data[0]));
