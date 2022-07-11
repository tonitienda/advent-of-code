const { getFileLines } = require('../input-tools')

const specialChars = { 1: 2, 4: 4, 7: 3, 8: 7 }

const main = () => {
    const segments = getFileLines('day-8/input.txt').flatMap(l => l.substr(l.indexOf('|') + 2).split(' '))

    const lengths = Object.values(specialChars)

    console.log(segments)
    console.log(lengths)
    console.log(segments.filter(s => lengths.includes(s.length)).length)

}

main()