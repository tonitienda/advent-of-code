const { getFileLines } = require('../input-tools')

const isLowPoint = (heights, i, j) => {
    const n = heights[i][j]

    if (i === 2 && j === 2) {
        console.log(n, heights[i][j - 1], heights[i][j + 1], heights[i - 1][j], heights[i + 1][j])
    }

    if (j > 0 && heights[i][j - 1] <= n) {
        return false
    }

    if (j < heights[i].length - 1 && heights[i][j + 1] <= n) {
        return false
    }

    if (i > 0 && heights[i - 1][j] <= n) {
        return false
    }


    if (i < heights.length - 1 && heights[i + 1][j] <= n) {
        return false
    }

    return true

}

function findBasin(heights, row, column, visited = []) {
    // Check if the point is already visited
    if (visited.some(b => b.row === row && b.column === column) || heights[row][column] === 9) {
        return []
    }
    const n = heights[row][column]
    visited.push({row, column})
    let basinPoints = [{ row, column }]

    if (column > 0 && heights[row][column - 1] > n) {
        basinPoints = basinPoints.concat(findBasin(heights, row, column - 1, visited))
    }

    if (column < heights[row].length - 1 && heights[row][column + 1] > n) {
        basinPoints = basinPoints.concat(findBasin(heights, row, column + 1, visited))
    }

    if (row > 0 && heights[row - 1][column] > n) {
        basinPoints = basinPoints.concat(findBasin(heights, row - 1, column, visited))
    }


    if (row < heights.length - 1 && heights[row + 1][column] > n) {
        basinPoints = basinPoints.concat(findBasin(heights, row + 1, column, visited))
    }

    return basinPoints.filter((b1, index, self) => self.findIndex((b2) => b1.row === b2.row && b1.column === b2.column) === index)
}


const main = () => {
    const heights = getFileLines('day-9/input.txt').map(l => l.split("").map(Number))

    let basins = []

    // Brute force
    for (let i = 0; i < heights.length; i++) {
        for (let j = 0; j < heights[i].length; j++) {
            if (isLowPoint(heights, i, j)) {
                basins.push(findBasin(heights, i, j))
            }
        }
    }

    // console.log(heights)
   // console.log(lowPoints)
    //console.log(lowPoints.map(n => n+1).reduce((sum, r) => sum + r, 0))
    const basinLengths = basins.map(b => b.length).sort((a,b) => b - a).filter((_, idx) => idx < 3)

    console.log(basinLengths.reduce((mul, l) => l*mul, 1))




    // for(let basin of basins) {
    //     console.log(`Basin`)
    //     console.log()
    //     for(let i = 0; i < heights.length ; i++){
    //         for(let j = 0; j<heights[i].length; j++) {
    //             if(basin.some(b => b.row === i && b.column === j)) {
    //                 process.stdout.write(`${heights[i][j]}`)
    //             } else {
    //                 process.stdout.write(` `)
    //             }
    //         }
    //         console.log()
    //     }
    //     console.log()
    // }

}

main()