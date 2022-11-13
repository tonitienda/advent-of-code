const getStringInBlocks = (str, blocksize) => {
  const arr = [];
  for (let i = 0; i < str.length; i++) {
    arr.push(str.substr(i, blocksize));
  }

  return arr.filter((a) => a.length === blocksize);
};

module.exports = {
  getStringInBlocks,
};
