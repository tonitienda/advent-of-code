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
    const size = points.length * 5 

    return [
        [i - 1, j],
        [i, j + 1],
        [i + 1, j],
        [i, j - 1]
    ].filter(p => p[0] >= 0 && p[0] < size && p[1] >= 0 && p[1] < size)

}

const findLowestNode = (distances, path) => {
    let lowestDistance = Infinity
    let node = null

    // console.log()
    // console.log(distances)
    for(let i =0; i<distances.length; i++){
        for(let j = 0; j< distances[i].length; j++) {
            const dist = getDistance(distances, i, j)
            // console.log([i,j], dist)
            if(dist < lowestDistance && !path[`${i}-${j}`]) {
                lowestDistance = dist
                node = [i,j] 

                // There is no lowest distance than 1
                if(lowestDistance <= 1) {
                    return node
                }
            }
        }
    }

    // console.log(node)
    // console.log()
    return node
}

const costsCache = []
const getCost = (points, i, j) => {
    if(!costsCache[i]) {
        costsCache[i]=[]
    }
    if(!costsCache[i][j]) {
        const adder1 = Math.floor(i / points.length)
        const adder2 = Math.floor(j / points.length)

        const cost = points[i % points.length][j % points.length]

        const sum =(cost + adder1 + adder2)
        const newCost = sum <= 9 ? sum : (sum % 10) + 1

        costsCache[i][j] = newCost
    }


    return costsCache[i][j]
}

const getDistance = (distances, i, j) => {

    // console.log('getDistance')
    // console.log(distances)
    // console.log(i,j)

    if(!distances[i] || distances[i][j] === undefined) {
        // console.log(Infinity)
        return Infinity
    }

    // console.log(distances[i][j])
    // console.log()
    return distances[i][j]
}

const setDistance = (distances, i, j, distance) => {
    if(!distances[i]) {
        distances[i]=[]
    }

    // console.log(`Set Distance`)
    // console.log(i,j, distance)
    // console.log()
    distances[i][j] = distance
}

const printDistances = (distances) => {
    console.log()
    for(var i=0; i< distances.length; i++){
        console.log()
        for(var j=0; j<distances[i].length; j++) {
            if(distances[i][j] >= 0) {
                process.stdout.write((distances[i][j]+"    ").substring(0,4))
            } else {
                process.stdout.write("    ")
            }
        }
    }
    console.log()
}


const printPoints = (points) => {
    console.log()
    for(var i=0; i< points.length * 5; i++){
        console.log()
        for(var j=0; j<points.length * 5; j++) {
            const cost = getCost(points, i, j)
            if(cost < Infinity) {
                process.stdout.write((cost+" ").substring(0,2))
            } else {
                process.stdout.write(" ")
            }
        }
    }
    console.log()
}



const dijkstra = (distances, points, destPoint) => {
    const visited = {}
    let found = false
    while(!found) {
        const node = findLowestNode(distances, visited)
        // console.log(`NODE`)
        // console.log(node)
        const id = `${node[0]}-${node[1]}`
        visited[id] = true

        console.log(`Visiting: `, node)

        if(node[0] === destPoint[0] && node[1] === destPoint[1] ) {
            return visited
        }

        const parentDistance = getDistance(distances, node[0], node[1])

        const adjacent = getCandidates(node, points)

        for(let [i,j] of adjacent) {
            const cost = getCost(points, i,j)
            const currentDistance = getDistance(distances, i, j)
            const finalDistance = Math.min(currentDistance, parentDistance + cost )

            // if(( i === 0 && j === 1) || (i === 1 && j === 0)) {
            //     console.log([i,j], {cost, currentDistance, parentDistance, finalDistance})
            // }
            

            // console.log(`Cost: `, cost, `currentDistance`, currentDistance, 'parentDistance', parentDistance)

            setDistance(distances, i, j,finalDistance)

        }
    }
}


const points = getFileLines('day-15/input.txt').map(l => l.split("").map(Number))
//const distances = createDistToStart(points)

// The matrix is a square
const last = (points.length * 5) -1
const destNode = [last, last]
//const destNode = [points.length -1, points.length -1]

const distances = []
distances[0] = [0]

const visited = dijkstra(distances, points, destNode)

const limit =  points.length * 5
console.log(`${limit}x${limit} (${limit*limit} cells)`)
//console.log(Object.keys(visited).map(k => k.split('-').map(Number).map(([i,j]) => points[i][j])))
//console.log(Object.keys(visited))
console.log(distances[destNode[0]][destNode[1]])

// printDistances(distances)


//printPoints(points)


