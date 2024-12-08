const fs = require("fs");

const INPUT_FILE = "./2024/day-8/input.txt";
const TEST_FILE = "./2024/day-8/test4.txt";

const content = fs
  .readFileSync(INPUT_FILE)
  .toString()
  .split("\n")
  .map((l) => l.split(""));

const antenaPositions = {};

console.log(content);
for (let row = 0; row < content.length; row++) {
  for (let col = 0; col < content[0].length; col++) {
    const c = content[row][col];

    if (c != ".") {
      if (!antenaPositions[c]) {
        antenaPositions[c] = [];
      }
      antenaPositions[c].push([row, col]);
    }
  }
}
console.log(antenaPositions);

const antinodes = [];
for (let points of Object.values(antenaPositions)) {
  for (let i = 0; i < points.length - 1; i++) {
    for (let j = i + 1; j < points.length; j++) {
      const pointA = points[i];
      const pointB = points[j];

      antinodes.push(pointA);
      antinodes.push(pointB);

      const v = [pointA[0] - pointB[0], pointA[1] - pointB[1]];

      let antinode1 = pointA;
      while (true) {
        antinode1 = [antinode1[0] + v[0], antinode1[1] + v[1]];

        if (
          antinode1[0] >= 0 &&
          antinode1[0] < content.length &&
          antinode1[1] >= 0 &&
          antinode1[1] < content[0].length
        ) {
          antinodes.push(antinode1);
        } else {
          break;
        }
      }

      let antinode2 = pointB;
      while (true) {
        antinode2 = [antinode2[0] - v[0], antinode2[1] - v[1]];
        if (
          antinode2[0] >= 0 &&
          antinode2[0] < content.length &&
          antinode2[1] >= 0 &&
          antinode2[1] < content[0].length
        ) {
          antinodes.push(antinode2);
        } else {
          break;
        }
      }
    }
  }
}

console.log("Antinodes:", antinodes);

console.log("Antinodes:", antinodes.length);

console.log(
  "Unique Antinodes:",
  antinodes.filter(
    (a, idx) =>
      antinodes.findIndex((a2) => a2[0] === a[0] && a2[1] === a[1]) == idx
  ).length
);
