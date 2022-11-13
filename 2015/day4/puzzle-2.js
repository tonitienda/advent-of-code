var crypto = require("crypto");

const md5 = (data) => crypto.createHash("md5").update(data).digest("hex");
const secretKey = "iwrupvqb";

const getFirstWithConditions = (key) => {
  let i = 1;

  while (true) {
    let hash = md5(key + i);

    if (hash.startsWith("000000")) {
      return i;
    }
    i++;
  }
};

const result = getFirstWithConditions(secretKey);

console.log(result);
