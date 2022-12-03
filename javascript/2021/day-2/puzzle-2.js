const { getFileLines } = require('../input-tools')

function main () {
    const {position, depth, aim} = getFileLines("day-2/input.txt").map(l => l.split(' '))
    .map(([direction, count]) => [direction, Number(count)])
    .reduce(({position, depth, aim}, [dir, x]) => {
        if(dir === 'forward') {
            return {
                position: position + x,
                depth: depth + aim * x,
                aim
            }
        }

        if(dir === 'down') {
            return {
                position,
                depth,
                aim: aim + x
            }
        }

        if(dir === 'up') {
            return {
                position,
                depth,
                aim: aim - x
            }
        }
    }, {position: 0, depth: 0, aim: 0})
   
    console.log({position, depth})
    console.log(position * depth)


}

main()