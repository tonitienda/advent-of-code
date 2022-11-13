const file = require("../../utils/file");
const path = require("path");
const arr = require("../../utils/array-extensions");

const contents = file.getChar2DArray(
  path.join(__dirname, "input.txt"),
  "\n",
  ""
);

const forbiddenStrings = ["ab", "cd", "pq", "xy"];

const vowels = ["a", "e", "i", "o", "u"];
const hasThreeVowels = (array) => arr.intersect(array, vowels).length >= 3;

const hasDoubleLetters = (array) => {
  for (let i = 0; i < array.length - 1; i++) {
    if (array[i] === array[i + 1]) {
      return true;
    }
  }
  return false;
};

const hasForbiddenStrings = (array) => {
  for (let i = 0; i < array.length - 1; i++) {
    let seq = array[i] + array[i + 1];

    if (forbiddenStrings.indexOf(seq) > -1) {
      return true;
    }
  }
  return false;
};

const isNice = (arr) =>
  hasThreeVowels(arr) && hasDoubleLetters(arr) && !hasForbiddenStrings(arr);

const areNiceContents = contents.map(isNice);

console.log(areNiceContents.filter((b) => b).length);
