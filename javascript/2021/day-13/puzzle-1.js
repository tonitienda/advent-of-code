const { getFileContent } = require('../input-tools')


const areUnique = (p, idx, self) => self.findIndex((p2) => p2[0] === p[0] && p2[1] === p[1]) === idx

const verticalFold = (hLine) => (board) => {
    return board.map(p => {
        if(p[1] > hLine) {
            return [p[0], hLine - (p[1] - hLine)]
        }
        return p
    }).filter(areUnique)

}


const horizontalFold = (vLine) => (board) => {
    return board.map(p => {
        if(p[0] > vLine) {
            return [vLine - (p[0] - vLine), p[1]]
        }
        return p
    }).filter(areUnique)

}

const printBoard = (board) => {
    const [maxCol, maxRow] = board.reduce((max, point) => [Math.max(max[0], point[0]), Math.max(max[1], point[1])])

    for (let j = 0; j <= maxRow; j++) {
        for (let i = 0; i <= maxCol; i++) {
            if (board.some(p => p[0] === i && p[1] === j)) {
                process.stdout.write('#')
            } else {
                process.stdout.write('.')
            }
        }
        console.log()
    }

}

const isLowerCase = (str) => str.toLowerCase() === str

const calculateRoutes = (graphDef, currentNode = 'start', smallVisited = ['start']) => currentNode === 'end' ? ['end'] : graphDef[currentNode].filter(n => !smallVisited.includes(n)).flatMap(n => calculateRoutes(graphDef, n, isLowerCase(n) ? smallVisited.concat(n) : smallVisited).map(r => [currentNode, ...r]))


const [pointsPart, instrPart] = getFileContent('day-13/input.txt', '\r\n\r\n')

let board = pointsPart.split('\r\n').map(l => l.split(',').map(Number))
let instructions = instrPart.split('\r\n').map(l => l.replace('fold along ', '')).map(i => i.startsWith('y=') ? verticalFold(Number(i.replace('y=',''))) :horizontalFold(Number(i.replace('x=', ''))))

    board = instructions[0](board)


//printBoard(board)
console.log(board.length)



