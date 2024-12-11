const fs = require("fs");

const INPUT_FILE = "./2024/day-10/input.txt";
const TEST_FILE = "./2024/day-10/test.txt";
const TEST2_FILE = "./2024/day-10/test2.txt";

const content = fs
  .readFileSync(TEST2_FILE)
  .toString()
  .split("\n")
  .map((l) => l.split("").map(Number));

console.log(content);

const getStartPoints = (matrix) => {
  const startPoints = [];

  for (let row = 0; row < matrix.length; row++) {
    for (let col = 0; col < matrix[row].length; col++) {
      if (matrix[row][col] === 0) {
        startPoints.push([row, col]);
      }
    }
  }

  return startPoints;
};

const EndPoint = 9;
const MaxSlope = 1;

const getNeighboursByContent = (content) => (row, col, maxSlope) => {
  const neighbours = [];

  const currentValue = content[row][col];
  if (content[row - 1] && content[row - 1][col] - maxSlope === currentValue) {
    neighbours.push([row - 1, col]);
  }

  if (content[row + 1] && content[row + 1][col] - maxSlope === currentValue) {
    neighbours.push([row + 1, col]);
  }

  if (
    content[row][col - 1] &&
    content[row][col - 1] - maxSlope === currentValue
  ) {
    neighbours.push([row, col - 1]);
  }

  if (
    content[row][col + 1] &&
    content[row][col + 1] - maxSlope === currentValue
  ) {
    neighbours.push([row, col + 1]);
  }

  return neighbours;
};

const getNeighbours = getNeighboursByContent(content);
const startPoints = getStartPoints(content);

const findPathByContent =
  (content) =>
  (currentPoint, endFound = []) => {
    const [row, col] = currentPoint;

    if (
      content[row][col] === EndPoint &&
      endFound.indexOf(([r, c]) => r === row && c === col) === -1
    ) {
      return [[row, col]];
    }

    const neighbours = getNeighbours(row, col, MaxSlope);
    console.log([row, col], "neighbours", neighbours);

    if (!neighbours) {
      return null;
    }

    let endPoints = [];

    for (let neighbour of neighbours) {
      const endpoints = findPath(neighbour, endPoints);

      if (endpoints) {
        endPoints = [...endPoints, ...endpoints];
      }
    }

    return endPoints.filter(
      ([r1, c1], idx) =>
        endPoints.findIndex(([r2, c2]) => r1 === r2 && c1 === c2) === idx
    );
  };

const findPath = findPathByContent(content);

const findPaths = (content) => {
  let allPaths = [];
  for (let startPoint of startPoints) {
    console.log("Start point:", startPoint);
    allPaths.push(findPath(startPoint));
  }

  return allPaths;
};

const paths = findPaths(content);

console.log("Trailheads = ", paths.length);

console.log(
  "Result = ",
  paths.map((t) => t.length).reduce((acc, n) => acc + n)
);
