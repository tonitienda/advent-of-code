const fs = require('fs');
const path = require('path')

const getFileContent = (localpath, splitBy) => fs.readFileSync(path.join(__dirname, localpath)).toString().split(splitBy)
const getFileLines = (localpath) => getFileContent(localpath, '\r\n')

const getNumbersFromFile = (localpath, splitBy = '\r\n') => getFileContent(localpath, splitBy).map(Number)

const printMatrix = (matrix, separator="") => {
    for(let i = 0; i < matrix.length; i++) {
        for(let j=0; j < matrix[i].length; j++) {
            process.stdout.write(`${matrix[i][j] === Infinity ? "" : matrix[i][j]}${separator}`)
        }
        console.log()
    }
}

module.exports = {
    getNumbersFromFile,
    getFileLines,
    printMatrix,
    getFileContent
}