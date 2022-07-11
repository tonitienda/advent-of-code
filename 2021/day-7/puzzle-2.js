const { getNumbersFromFile } = require('../input-tools')

const cache = {}
const memoize = (fn) => (n) => {
    if(!cache[n]) {
        cache[n] = fn(n)
    }

    return cache[n]
}




const cost = memoize((n) => {
    if(n <= 1) {
        return n
    }

    return n + cost( n - 1)
})

const main = () => {
    const crabPositions = getNumbersFromFile('day-7/input.txt', ',')

    const min = crabPositions.reduce((min, n) => n < min ? n : min)
    const max = crabPositions.reduce((max, n) => n > max ? n : max)

    // console.log(crabPositions)
    // console.log(min,max)

    let minFuel = Infinity
    let pos = -1
    for(let i = min; i<=max; i++) {
        const fuel = crabPositions.map(n => Math.abs(n - i)).reduce((sum, n) => cost(n) + sum, 0)

        if(fuel < minFuel) {
            minFuel = fuel
            pos = i
        }
    }

    console.log({minFuel,pos})


}

main()