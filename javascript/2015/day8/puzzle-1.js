const file = require("../../utils/file");
const path = require("path");

const formatLine = (line) =>
  line
    .slice(1, -1)
    .replace(/\\\\/g, "_")
    .replace(/\\x[0-9a-f][0-9a-f]/g, "0")
    .replace(/\\"/g, "=");

const contents = file.getLines(path.join(__dirname, "input.txt"), "\n", "");

console.log(contents);

const totalChars = contents.reduce((total, line) => total + line.length, 0);

console.log(totalChars);

let parsedLines = contents.map(formatLine);

console.log(parsedLines);

const totalMemoryChars = parsedLines.reduce(
  (total, line) => total + line.length,
  0
);
console.log(totalMemoryChars);

console.log(
  totalChars,
  "-",
  totalMemoryChars,
  "=",
  totalChars - totalMemoryChars
);
