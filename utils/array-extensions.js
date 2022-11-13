const sum = (arr) => arr.reduce((total, item) => total + item, 0);
const mult = (arr) => arr.reduce((total, item) => total * item, 1);

const sort = (arr) => arr.sort((a, b) => (a < b ? -1 : 1));
const intersect = (arr1, arr2) => arr1.filter((a) => arr2.indexOf(a) > -1);

const initArray = (rows, columns, initialValue) => {
  const arr = [];

  for (let i = 0; i < rows; i++) {
    arr[i] = [];
    for (let j = 0; j < columns; j++) {
      arr[i][j] = initialValue;
    }
  }
  return arr;
};

module.exports = {
  sum,
  mult,
  sort,
  intersect,
  initArray,
};
