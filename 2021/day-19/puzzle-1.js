const { getFileLines } = require("../input-tools");

const lines = getFileLines("day-19/test.txt");

let NoRotation = [
  [1, 0, 0],
  [0, 1, 0],
  [0, 0, 1],
];

let RotationX = (angle) =>
  angle === 0
    ? NoRotation
    : [
        [1, 0, 0],
        [0, Math.cos(angle), -Math.sin(angle)],
        [0, Math.sin(angle), Math.cos(angle)],
      ];

let RotationY = (angle) =>
  angle === 0
    ? NoRotation
    : [
        [Math.cos(angle), 0, Math.sin(angle)],
        [0, 1, 0],
        [-Math.sin(angle), 0, Math.cos(angle)],
      ];
let RotationZ = (angle) =>
  angle === 0
    ? NoRotation
    : [
        [Math.cos(angle), -Math.sin(angle), 0],
        [Math.sin(angle), Math.cos(angle), 0],
        [1, 0, 0],
      ];

let RotationPossibilities = [
  [0, 0, 0],
  [0, 0, 90],
  [0, 0, 180],
  [0, 0, 270],
  [0, 90, 0],
  [0, 90, 90],
  [0, 90, 180],
  [0, 90, 270],
  [0, 180, 0],
  [0, 180, 90],
  [0, 180, 180],
  [0, 180, 270],
  [0, 270, 0],
  [0, 270, 90],
  [0, 270, 180],
  [0, 270, 270], 
  [90, 0, 90],
  [90, 0, 180],
  [90, 0, 270],
  [90, 90, 0],
  [90, 180, 0],
  [90, 270, 0],
  [180, 0, 90],
  [180, 0, 180],
  [180, 0, 270],
  [180, 90, 0],
  [180, 180, 0],
  [180, 270, 0],
  [270, 0, 90],
  [270, 0, 180],
  [270, 0, 270],
  [270, 90, 0],
  [270, 180, 0],
  [270, 270, 0],
];

const Mult = (m1, m2) => m1.map((r, i) => r.map((n, j) => n * m2[i][j]) )

let RotationMatrices = RotationPossibilities.map(
    r => [RotationX(r[0]), RotationY(r[1]), RotationZ(r[2])].reduce((acc, m) => Mult(acc, m))
)

console.log(RotationMatrices, RotationMatrices.length)

let currentScanner = -1;
const scannerData = lines.reduce((acc, line) => {
  
  if (line.startsWith("--- scanner ")) {
    currentScanner = Number(
      line.replace("--- scanner ", "").replace(" ---", "")
    );
    acc[currentScanner] = [];
    return acc;
  }

  acc[currentScanner].push(line.split(",").map(Number));
  return acc;
}, []);
