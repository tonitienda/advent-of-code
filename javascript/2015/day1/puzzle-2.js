const file = require("../../utils/file");
const path = require("path");

const charMod = {
  "(": 1,
  ")": -1,
};
const contents = file.getCharArray(path.join(__dirname, "input.txt"), "");

const firstBasement = (contents, charMod) => {
  let floor = 0;
  for (let i = 0; i < contents.length; i++) {
    floor += charMod[contents[i]];

    if (floor === -1) {
      return i + 1;
    }
  }
};

const result = firstBasement(contents, charMod);

console.log("First basement:", result);
