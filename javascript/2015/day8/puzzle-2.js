const file = require("../../utils/file");
const path = require("path");

const formatLine = (line) =>
  '"' + line.replace(/\\/g, "\\\\").replace(/"/g, '\\"') + '"';

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
//console.log(totalMemoryChars);

console.log(
  totalChars,
  "-",
  totalMemoryChars,
  "=",
  totalChars - totalMemoryChars
);
