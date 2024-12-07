const fs = require("fs");

const INPUT_FILE = "./2024/day-5/input.txt";
const TEST_FILE = "./2024/day-5/test.txt";

const content = fs.readFileSync(INPUT_FILE).toString();

console.log(content);

const [rulesStr, updatesStr] = content.split("\n\n");

// Create graph based on positions
const rules = rulesStr
  .split("\n")
  .map((l) => l.split("|"))
  .reduce((ruleSet, rule) => {
    const [origin, dest] = rule;
    if (!ruleSet[origin]) {
      ruleSet[origin] = [];
    }

    if (!ruleSet[dest]) {
      ruleSet[dest] = [origin];
    } else {
      ruleSet[dest].push(origin);
    }
    return ruleSet;
  }, {});

console.log(rules);

const isRightOrder2 = (rules) => (line) => {
  for (let i = 0; i < line.length - 1; i++) {
    const predecessors = rules[line[i]];

    for (let j = i + 1; j < line.length; j++) {
      const second = line[j];

      if (predecessors.indexOf(second) > -1) {
        return false;
      }
    }
  }

  return true;
};

const correctUpdates = (rules) => (line) => {
  // console.log();
  // console.log("rules:", rules);
  // console.log("Initial:", line);
  for (let i = 0; i < line.length - 1; i++) {
    let predecessors = rules[line[i]];

    for (let j = i + 1; j < line.length; j++) {
      const second = line[j];

      if (predecessors.indexOf(second) > -1) {
        // console.log("swapping", line[i], "<->", line[j]);
        const aux = line[i];
        line[i] = line[j];
        line[j] = aux;

        predecessors = rules[line[i]];
      }
    }
  }

  //console.log("Final:", line);
  return line;
};
const correctedUpdates = updatesStr
  .split("\n")
  .map((l) => l.split(","))
  .filter((l) => !isRightOrder2(rules)(l))
  .map(correctUpdates(rules));

console.log(correctedUpdates);

const result = correctedUpdates
  .map((u) => Number(u[Math.floor(u.length / 2)]))
  .reduce((sum, n) => sum + n, 0);

// console.log(correctUpdates);
console.log(result);
