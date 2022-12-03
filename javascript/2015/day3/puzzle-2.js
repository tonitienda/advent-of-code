const file = require("../../utils/file");
const path = require("path");

const contents = file.getCharArray(path.join(__dirname, "input.txt"), "");

console.log(contents);

const santaMoves = contents.filter((c, idx) => idx % 2 === 0);
const santaRobotMoves = contents.filter((c, idx) => idx % 2 === 1);

const calculateUniqueHousesVisited = (contents) => {
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

  const uniqueHouses = visitedHouses.filter(
    (h, index, self) =>
      self.findIndex((h2) => h2[0] === h[0] && h2[1] === h[1]) === index
  );

  return uniqueHouses.length;
};

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

const visitedHousesBySanta = calculateHousesVisited(santaMoves);
const visitedHousesByRobot = calculateHousesVisited(santaRobotMoves);

const uniqueHouses = [...visitedHousesBySanta, ...visitedHousesByRobot].filter(
  (h, index, self) =>
    self.findIndex((h2) => h2[0] === h[0] && h2[1] === h[1]) === index
);

console.log(uniqueHouses.length);
