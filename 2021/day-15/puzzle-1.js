const { getFileLines, printMatrix } = require('../input-tools')

const createDistToStart = (points) => {
    const distances = []

    for (let i = 0; i < points.length; i++) {
        distances[i] = []
        for (let j = 0; j < points[i].length; j++) {
            distances[i][j] = Infinity
        }
    }

    return distances
}



const getCandidates = ([i, j], points) => {
    return [
        [i - 1, j],
        [i, j + 1],
        [i + 1, j],
        [i, j - 1]
    ].filter(p => p[0] >= 0 && p[0] < points.length && p[1] >= 0 && p[1] < points[points.length - 1].length)

}

const findLowestNode = (distances, path) => {
    let lowestDistance = Infinity
    let node = null

    for(let i =0; i<distances.length; i++){
        for(let j = 0; j< distances[i].length; j++) {
            if(distances[i][j] !== Infinity && distances[i][j] < lowestDistance && !path[`${i}-${j}`]) {
                lowestDistance = distances[i][j]
                node = [i,j] 

                // There is no lowest distance than 1
                if(lowestDistance <= 1) {
                    return node
                }
            }
        }
    }

    return node
}

const dijkstra = (distances, points, destPoint) => {
    const path = {}
    let found = false
    while(!found) {
        const node = findLowestNode(distances, path)
        const id = `${node[0]}-${node[1]}`
        path[id] = true

        if(node[0] === destPoint[0] && node[1] === destPoint[1] ) {
            return path
        }

        const parentDistance = distances[node[0]][node[1]]

        const adjacent = getCandidates(node, points)

        for(let [i,j] of adjacent) {
            const cost = points[i][j]
            const currentDistance = distances[i][j]
            distances[i][j] = Math.min(currentDistance, parentDistance + cost )

        }
    }
}


const points = getFileLines('day-15/input.txt').map(l => l.split("").map(Number))
const distances = createDistToStart(points)

// The matrix is a square
const destNode = [points.length-1, points.length-1]

// Start node
distances[0][0] = 0

const path = dijkstra(distances, points, destNode)
console.log()
console.log(Object.keys(path))
console.log(distances[destNode[0]][destNode[1]])


