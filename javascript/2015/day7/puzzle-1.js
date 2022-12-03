const file = require("../../utils/file");
const path = require("path");

const contents = file.getLines(path.join(__dirname, "input.txt"), "\n", "");

const getValue = (v, db) => (v in db ? db[v] : Number(v));

const assignment = ([source, dest], db) => {
  const value = getValue(source, db);
  if (Number.isNaN(value)) {
    return false;
  }

  db[dest] = value;
  return true;
};

const not = ([_, source, dest], db) => {
  const value = getValue(source, db);
  if (Number.isNaN(value)) {
    return false;
  }

  db[dest] = 65536 + ~value;
  return true;
};

const and = ([a, _, b, dest], db) => {
  const val_a = getValue(a, db);
  const val_b = getValue(b, db);

  if (Number.isNaN(val_a) || Number.isNaN(val_b)) {
    return false;
  }

  db[dest] = val_a & val_b;
  return true;
};

const or = ([a, _, b, dest], db) => {
  const val_a = getValue(a, db);
  const val_b = getValue(b, db);

  if (Number.isNaN(val_a) || Number.isNaN(val_b)) {
    return false;
  }

  db[dest] = val_a | val_b;
  return true;
};

const lshift = ([a, _, b, dest], db) => {
  const val_a = getValue(a, db);
  const val_b = getValue(b, db);

  if (Number.isNaN(val_a) || Number.isNaN(val_b)) {
    return false;
  }

  db[dest] = val_a << val_b;
  return true;
};

const rshift = ([a, _, b, dest], db) => {
  const val_a = getValue(a, db);
  const val_b = getValue(b, db);

  if (Number.isNaN(val_a) || Number.isNaN(val_b)) {
    return false;
  }

  db[dest] = val_a >> val_b;
  return true;
};

const run = (instruction, db) => {
  if (instruction.length === 2) {
    return assignment(instruction, db);
  }

  if (instruction[0] === "NOT") {
    return not(instruction, db);
  }

  if (instruction[1] === "AND") {
    return and(instruction, db);
  }

  if (instruction[1] === "OR") {
    return or(instruction, db);
  }

  if (instruction[1] === "LSHIFT") {
    return lshift(instruction, db);
  }

  if (instruction[1] === "RSHIFT") {
    return rshift(instruction, db);
  }

  return false;
};

let instructions = contents
  .map((l) => l.split(" -> "))
  .map(([instr, dest]) => [...instr.split(" "), dest]);

let db = {};

while (instructions.length > 0) {
  let pending = [];
  for (let instruction of instructions) {
    let wasExecuted = run(instruction, db);

    if (!wasExecuted) {
      pending.push(instruction);
    }
  }

  instructions = pending;
}

console.log(db);
