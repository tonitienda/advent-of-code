const fs = require("fs");

const INPUT_FILE = "./2024/day-6/input.txt";
const TEST_FILE = "./2024/day-6/test.txt";

const content = fs.readFileSync(TEST_FILE).toString();

console.log(content);
const contentArr = content.split("\n").map((l) => l.split(""));

const analyze = (arr) => {
  let initialPos = null;
  const obstacles = {};

  for (let row = 0; row < arr.length; row++) {
    for (let col = 0; col < arr[row].length; col++) {
      if (arr[row][col] === "^") {
        initialPos = [row, col];
      }

      if (arr[row][col] === "#") {
        if (!obstacles[row]) {
          obstacles[row] = {};
        }

        obstacles[row][col] = true;
      }
    }
  }

  return [initialPos, obstacles];
};

const UP = [-1, 0];
const DOWN = [1, 0];
const LEFT = [0, -1];
const RIGHT = [0, 1];

const directions = [UP, RIGHT, DOWN, LEFT];

const [initialPos, obstacles] = analyze(contentArr);

console.log(initialPos);
console.log(obstacles);

const walkMesh = (
  initialPos,
  initialDirection,
  obstacles,
  numRows,
  numCols
) => {
  const obstaclesFound = [];
  const walkedCells = {};

  walkedCells[initialPos[0]] = {};
  walkedCells[initialPos[0]][initialPos[1]] = true;

  let position = initialPos;
  let direction = initialDirection;

  while (true) {
    const nextRow = position[0] + directions[direction][0];
    const nextCol = position[1] + directions[direction][1];

    if (
      nextRow < 0 ||
      nextRow === numRows ||
      nextCol < 0 ||
      nextCol === numCols
    ) {
      return [walkedCells, obstaclesFound];
    }

    if (obstacles[nextRow] && obstacles[nextRow][nextCol]) {
      console.log("Obstacle found at ", nextRow, nextCol);
      obstaclesFound.push([nextRow, nextCol]);

      direction = (direction + 1) % directions.length;
    } else {
      position[0] = nextRow;
      position[1] = nextCol;

      if (!walkedCells[nextRow]) {
        walkedCells[nextRow] = {};
      }

      walkedCells[nextRow][nextCol] = true;
    }
  }
};

const [walkedCells, obstaclesFound] = walkMesh(
  initialPos,
  0,
  obstacles,
  contentArr.length,
  contentArr[0].length
);

console.log("====================");
console.log(obstaclesFound);

const findCloseLoopCandidates = (obstacles) => {
  //for (let i = 0; i < obstacles.length - 1; i++) {
  const i = 0;
  const direction = directions[i];
  const first = obstacles[i];
  const second = obstacles[i + 1];
  const third = obstacles[i + 2];

  // }
};
