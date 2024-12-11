const fs = require("fs");

const INPUT_FILE = "./2024/day-9/input.txt";
const TEST_FILE = "./2024/day-9/test.txt";

const content = fs.readFileSync(INPUT_FILE).toString().split("").map(Number);

console.log(content);

const files = content.filter((_, idx) => idx % 2 == 0);
//.map((f, idx) => new Array(f).fill(idx));
const spaces = content.filter((_, idx) => idx % 2 == 1);

console.log("FILES");
console.log(files);

console.log("SPACES");
console.log(spaces);

let result = 0;
let str = "";

let idx = 0;
let finalIdx = files.length - 1;
let expandedIdx = 0;

while (true) {
  const file = files[idx];

  for (let i = 0; i < file; i++) {
    str += idx;
    result += expandedIdx * idx;
    expandedIdx++;
    files[idx]--;
  }

  while (spaces[idx] > 0) {
    if (files[finalIdx] === 0) {
      finalIdx--;
      if (files[finalIdx] === 0) {
        console.log("break with str=", str);
        break;
      }
    }

    str += finalIdx;
    result += expandedIdx * finalIdx;
    expandedIdx++;

    files[finalIdx]--;
    spaces[idx]--;
  }
  idx++;

  if (finalIdx < idx) {
    console.log("Files:", files);
    console.log("Spaces:", spaces);
    console.log("finalIdx:", finalIdx);
    console.log("idx:", idx);
    break;
  }
}

console.log();
console.log(str);
console.log();

console.log("Result=", result);
