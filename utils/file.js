const fs = require("fs");

const getText = (path) => fs.readFileSync(path).toString();

const getCharArray = (path, split = ",") => getText(path).split(split);

const getNumber2DArray = (path, split1 = "\n", split2 = ",") =>
  getText(path)
    .split(split1)
    .map((r) => r.split(split2).map(Number));

const getChar2DArray = (path, split1 = "\n", split2 = ",") =>
  getText(path)
    .split(split1)
    .map((r) => r.split(split2));

const getLines = (path, split = "\n") => getText(path).split(split);

const getWords = (path, split = "\n") =>
  getText(path)
    .split(split)
    .map((l) => l.split(" "));

module.exports = {
  getCharArray,
  getNumber2DArray,
  getChar2DArray,
  getLines,
  getWords,
};
