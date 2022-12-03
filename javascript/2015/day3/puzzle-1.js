const file = require("../../utils/file");
const path = require("path");

const contents = file.getCharArray(path.join(__dirname, "input.txt"), "");

console.log(contents);

const calculateHousesVisited = (contents) => {
  let currentHouse = [0, 0];
  let visitedHouses = [currentHouse];

  const moves = {
    "^": ([x, y]) => [x, y - 1],
    v: ([x, y]) => [x, y + 1],
    ">": ([x, y]) => [x + 1, y],
    "<": ([x, y]) => [x - 1, y],
  };

  for (let s of contents) {
    currentHouse = moves[s](currentHouse);
    visitedHouses.push(currentHouse);
  }

  return visitedHouses;
};

const visitedHouses = calculateHousesVisited(contents);
const uniqueHouses = visitedHouses.filter(
  (h, index, self) =>
    self.findIndex((h2) => h2[0] === h[0] && h2[1] === h[1]) === index
);

console.log(uniqueHouses.length);
