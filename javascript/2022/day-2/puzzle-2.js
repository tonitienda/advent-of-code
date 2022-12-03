const file = require("../../utils/file");
const { sum } = require("../../utils/array-extensions");
const path = require("path");

const data = file
  .getLines(path.join(__dirname, "input.txt"))
  .map((line) => line.split(" "));

const Rock_1 = "A";
const Paper_1 = "B";
const Scissors_1 = "C";

const Loose = "X";
const Draw = "Y";
const Win = "Z";

const Rock_Points = 1;
const Paper_Points = 2;
const Scissor_Points = 3;

const Loose_Points = 0;
const Draw_Points = 3;
const Win_Points = 6;

const points = {
  [Rock_1]: {
    [Draw]: Rock_Points + Draw_Points,
    [Win]: Paper_Points + Win_Points,
    [Loose]: Scissor_Points + Loose_Points,
  },
  [Paper_1]: {
    [Loose]: Rock_Points + Loose_Points,
    [Draw]: Paper_Points + Draw_Points,
    [Win]: Scissor_Points + Win_Points,
  },
  [Scissors_1]: {
    [Win]: Rock_Points + Win_Points,
    [Loose]: Paper_Points + Loose_Points,
    [Draw]: Scissor_Points + Draw_Points,
  },
};

const gamePoints = data.map((s) => points[s[0]][s[1]]);
console.log(sum(gamePoints));
