const file = require("../../utils/file");
const path = require("path");
const str = require("../../utils/string-extensions");

const contents = file.getLines(path.join(__dirname, "input.txt"), "\n", "");

console.log(contents);

const blocks = contents.map((c) => ({
  original: c,
  block3: str.getStringInBlocks(c, 3),
  block2: str.getStringInBlocks(c, 2),
}));

console.log(blocks);

const hasPairTwice = (data) => {
  const { block2 } = data;

  for (let i = 0; i < block2.length - 1; i++) {
    for (let j = i + 1; j < block2.length; j++) {
      if (block2[i] === block2[j]) {
        // Overlaps means that the later in the pair is repeated (e.g. aa)
        // and the next pair we are looking for is in the consecutive pair ( j = i + 1 )
        const overlaps = block2[i][0] === block2[i][1] && j === i + 1;
        if (!overlaps) return true;
      }
    }
  }
  return false;
};

const hasLetterBetween = (data) => {
  const { block3 } = data;

  for (let i = 0; i < block3.length; i++) {
    if (block3[i][0] === block3[i][2]) {
      return true;
    }
  }
  return false;
};

const withPairs = blocks.map((b) => ({
  ...b,
  pair: hasPairTwice(b),
  letterBetween: hasLetterBetween(b),
}));

console.log(withPairs);

console.log(withPairs.filter((w) => w.pair && w.letterBetween).length);
