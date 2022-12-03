const { getFileLines } = require('../input-tools')


const isLowerCase = (str) => str.toLowerCase() === str

const calculateRoutes = (graphDef, currentNode = 'start', smallVisited = ['start']) => currentNode === 'end' ? ['end'] : graphDef[currentNode].filter(n => !smallVisited.includes(n)).flatMap(n => calculateRoutes(graphDef, n, isLowerCase(n) ? smallVisited.concat(n) : smallVisited).map(r => [currentNode, ...r]))


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




