const { getFileLines } = require('../input-tools')

function main () {
    const {position, depth} = getFileLines("day-2/input.txt").map(l => l.split(' '))
    .map(([direction, count]) => [direction, Number(count)])
    .reduce(({position, depth}, [dir, x]) => {
        if(dir === 'forward') {
            return {
                position: position + x,
                depth
            }
        }

        if(dir === 'down') {
            return {
                position,
                depth: depth + x
            }
        }

        if(dir === 'up') {
            return {
                position,
                depth: depth - x
            }
        }
    }, {position: 0, depth: 0})
   
    console.log({position, depth})
    console.log(position * depth)


}

main()