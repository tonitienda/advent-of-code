const file = require("../../utils/file");
const path = require("path");
const arr = require("../../utils/array-extensions");

const coords = (s) => s.split(",").map(Number);

const contents = file.getLines(path.join(__dirname, "input.txt"), "\n", "");

const instructions = contents
  .map((c) =>
    c
      .replace(/turn on/g, "turn-on")
      .replace(/turn off/g, "turn-off")
      .split(" ")
  )
  .map(([inst, from, _, to]) => [inst, coords(from), coords(to)]);

const iterate = (from, to, arr, fn) => {
  for (let i = from[0]; i <= to[0]; i++) {
    for (let j = from[1]; j <= to[1]; j++) {
      arr[i][j] = fn(arr[i][j]);
    }
  }
  return arr;
};

const Handlers = {
  toggle: (from, to, arr) => iterate(from, to, arr, (v) => v + 2),
  "turn-on": (from, to, arr) => iterate(from, to, arr, (v) => v + 1),
  "turn-off": (from, to, arr) =>
    iterate(from, to, arr, (v) => Math.max(v - 1, 0)),
};

const execute = (matrixSize, instructions) => {
  const lights = arr.initArray(matrixSize[0], matrixSize[1], 0);

  const endResult = instructions.reduce(
    (currentState, [instr, from, to]) =>
      Handlers[instr](from, to, currentState),
    lights
  );

  return endResult;
};

const instr = [
  ["toggle", [0, 0], [9, 9]],
  ["turn-off", [4, 2], [7, 5]],
];

const result = execute([1000, 1000], instructions);

const intensity = (lights) =>
  lights.reduce((total, r) => arr.sum(r) + total, 0);

console.log(intensity(result));
