const { getFileLines, printMatrix } = require('../input-tools')

function increaseCell(energyLevels, row, col, flashedCells) {
    if (row < 0 || row > energyLevels.length - 1) {
        return
    }

    if (col < 0 || col > energyLevels[row].length - 1) {
        return
    }

    if(flashedCells.some(c => c.row === row && c.col === col)){
        return
    }

    energyLevels[row][col]++
}

function simulateFlashes(energyLevels, flashedCells = []) {
    let flashes = 0

    for (let row = 0; row < energyLevels.length; row++) {
        for (let col = 0; col < energyLevels[row].length; col++) {
            energyLevels[row][col]++
        }
    }

    let somethingHasFlashed = false

    do {
        somethingHasFlashed = false
        for (let row = 0; row < energyLevels.length; row++) {
            for (let col = 0; col < energyLevels[row].length; col++) {
                if (energyLevels[row][col] > 9 && !flashedCells.some(c => c.row === row && c.col === col)) {
                    // Trackcell that has flashed
                    flashedCells.push({ row, col })
                    somethingHasFlashed = true
                    energyLevels[row][col] = 0
                    flashes++

                    increaseCell(energyLevels, row - 1, col, flashedCells)
                    increaseCell(energyLevels, row - 1, col + 1, flashedCells)
                    increaseCell(energyLevels, row, col + 1, flashedCells)
                    increaseCell(energyLevels, row + 1, col + 1, flashedCells)
                    increaseCell(energyLevels, row + 1, col, flashedCells)
                    increaseCell(energyLevels, row + 1, col - 1, flashedCells)
                    increaseCell(energyLevels, row, col - 1, flashedCells)
                    increaseCell(energyLevels, row - 1, col - 1, flashedCells)
                }
            }
        }
    } while (somethingHasFlashed)

    return flashes
}

const energyLevels = getFileLines('day-11/input.txt').map(l => l.split("").map(Number))



printMatrix(energyLevels)

let totalFlashes = 0
for (let i = 0; i < 100; i++) {
    totalFlashes += simulateFlashes(energyLevels)
}   

printMatrix(energyLevels)

console.log(totalFlashes)

