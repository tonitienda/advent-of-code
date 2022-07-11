const { getFileLines } = require('../input-tools')


const isLowerCase = (str) => str.toLowerCase() === str

const canVisit = (node, smallVisited, twiceVisited) => {
    if(node === 'start') {
        return [false, twiceVisited]
    }
    if(!isLowerCase(node)) {
        return [true, twiceVisited]
    }

    if(smallVisited.includes(node)) {
        if(twiceVisited) {
            return [false, true]
        } else {
            return [true, true]
        }
        //return [!twiceVisited, true]
    }

    return [true, twiceVisited]

}

const calculateRoutes = (graphDef, currentNode = 'start', smallVisited = ['start'], twiceVisited = false) => {
    if(currentNode === 'end') {
        return['end']
    }

    let routes = []
    for(const node of graphDef[currentNode]) {

        const [canGo, newTwiceVisited] = canVisit(node, smallVisited, twiceVisited)
        if(canGo) {
            routes = routes.concat(calculateRoutes(graphDef, node, isLowerCase(node)? smallVisited.concat(node) : smallVisited, newTwiceVisited))
        }
    }

    return routes.map(r => ([currentNode, ...r]))
    
}

const graphDef = getFileLines('day-12/input.txt').map(l => l.split('-')).reduce((graph, line) => {
    const [from, to] = line
    if (!graph[from]) {
        graph[from] = []
    }

    if (!graph[to]) {
        graph[to] = []
    }

    graph[from].push(to)
    graph[to].push(from)

    return graph

}, {})

console.log('----')
const routes = calculateRoutes(graphDef)
//console.log(routes)

console.log(routes.length)




