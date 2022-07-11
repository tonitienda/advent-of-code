const { getFileLines } = require('../input-tools')

const main = () => {
const lines = getFileLines('day-5/input.txt').map(l => l.split(' -> ').map(c => c.split(',').map(Number)))

console.log(lines)


const points = {}

const addPoint = (points, [x,y]) => {
    const key = `${x}-${y}`
    if(!points[key]) {
        points[key]=1
    } else {
        points[key]++
    }
}

for(const line of lines) {
    const [[x1,y1],[x2,y2]] = line

    if(x1 === x2 ) {
        const start = Math.min(y1,y2)
        const end = Math.max(y1,y2)

        for(let y = start; y<= end; y++){
            addPoint(points, [x1, y])

        }
    }

    if(y1 === y2 ) {
        const start = Math.min(x1,x2)
        const end = Math.max(x1,x2)

        for(let x = start; x<= end; x++){
            addPoint(points, [x, y1])
        }
    }
}

// console.log(points)

const total = Object.values(points).filter(v => v > 1).length

console.log(total)
}

main()  