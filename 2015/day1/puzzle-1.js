const file = require("../../utils/file");
const path = require("path");

const charMod = {
  "(": 1,
  ")": -1,
};
const contents = file.getCharArray(path.join(__dirname, "input.txt"), "");

console.log(contents);
const floor = contents.reduce(
  (currentFloor, char) => currentFloor + charMod[char],
  0
);

console.log("Floor:", floor);
