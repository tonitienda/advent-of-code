const fs = require("fs");

const INPUT_FILE = "./2024/day-11/input.txt";
const TEST_FILE = "./2024/day-11/test.txt";

const content = fs.readFileSync(INPUT_FILE).toString().split(" ").map(Number);

console.log(content);

const blink = (n, pendingSteps) => {
  if (pendingSteps === 0) {
    return 1;
  }

  if (n === 0) {
    return blink(1, pendingSteps - 1);
  }

  // This is slow
  const s = "" + n;
  if (s.length % 2 === 0) {
    const s1 = s.substring(0, s.length / 2);
    const s2 = s.substring(s.length / 2);

    return (
      blink(Number(s1), pendingSteps - 1) + blink(Number(s2), pendingSteps - 1)
    );
  }

  return blink(n * 2024, pendingSteps - 1);
};

const totalStones = content
  .map((s) => blink(s, 25))
  .reduce((acc, i) => acc + i, 0);

console.log("Result:", totalStones);
