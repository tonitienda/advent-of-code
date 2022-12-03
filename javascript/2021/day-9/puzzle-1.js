const { getFileLines } = require('../input-tools')

const isLowPoint = (heights, i,j) => {
    const n = heights[i][j]

    if(i === 2 && j === 2) {
        console.log(n, heights[i][j-1], heights[i][j+1],heights[i-1][j], heights[i+1][j] )
    }

    if(j > 0 && heights[i][j-1] <= n){
        return false
    }

    if(j < heights[i].length - 1 && heights[i][j+1] <= n) {
        return false
    }
    
    if(i > 0 && heights[i-1][j] <= n) {
        return false
    }


    if(i < heights.length - 1 && heights[i+1][j] <= n) {
        return false
    }

    return true

}


const main = () => {
    const heights = getFileLines('day-9/input.txt').map(l => l.split("").map(Number))

    let lowPoints = []

    // Brute force
    for(let i = 0; i< heights.length; i++) {
        for(let j = 0; j < heights[i].length; j++) {
            if(isLowPoint(heights, i,j)) {
                console.log(i, j)
                lowPoints.push(heights[i][j])
            }
        }
    }
   // console.log(heights)
    console.log(lowPoints)
    console.log(lowPoints.map(n => n+1).reduce((sum, r) => sum + r, 0))

}

main()