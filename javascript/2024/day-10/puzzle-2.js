const fs = require("fs");

const INPUT_FILE = "./2024/day-10/input.txt";
const TEST_FILE = "./2024/day-10/test.txt";
const TEST2_FILE = "./2024/day-10/test2.txt";

const content = fs
  .readFileSync(INPUT_FILE)
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
  (currentPoint, walkedCells = []) => {
    const [row, col] = currentPoint;
    walkedCells.push(currentPoint);

    if (content[row][col] === EndPoint) {
      return [walkedCells];
    }

    const neighbours = getNeighbours(row, col, MaxSlope);

    if (!neighbours) {
      return null;
    }

    let allPaths = [];

    for (let neighbour of neighbours) {
      const paths = findPath(neighbour, walkedCells);

      if (paths) {
        allPaths = [...allPaths, ...paths];
      }
    }

    return allPaths;
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
