const { getFileLines } = require('../input-tools')


const lengthToNumber = { 2: 1, 4: 4, 3: 7, 7: 8 }


const letters = ["abcefg", "ab", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdef", "abcdfg"]

// All letters in w2 are found in w1
const contains = (w1, w2) => {
    return w2.split("").every(l => w1.indexOf(l) > -1)
}

// Some letters in w2 are not found in w1
const doesNotContain = (w1, w2) => {
    return w2.split("").some(l => w1.indexOf(l) === -1)
}


const areSame = (w1, w2) => {
    return w1.length === w2.length && contains(w1, w2)
}

const segmentsToNumbers = (segments) => {
    const one = segments.find(w => w.length === 2).split("").sort().join("")
    const seven = segments.find(w => w.length === 3).split("").sort().join("")
    const four = segments.find(w => w.length === 4).split("").sort().join("")
    const eight = segments.find(w => w.length === 7).split("").sort().join("")
    const three = segments.find(w => w.length === 5 && contains(w, seven)).split("").sort().join("")
    const nine = segments.find(w => w.length === 6 && contains(w, four)).split("").sort().join("")
    const zero = segments.find(w => w.length === 6 && contains(w, one) && doesNotContain(w, four)).split("").sort().join("")
    const six = segments.find(w => w.length === 6 && doesNotContain(w, one)).split("").sort().join("")
    const five = segments.find(w => w.length === 5 && contains(six, w)).split("").sort().join("")
    const two = segments.find(w => w.length === 5 && doesNotContain(nine, w)).split("").sort().join("")

    return [
        zero,
        one,
        two,
        three,
        four,
        five,
        six,
        seven,
        eight,
        nine
    ]

}

const main = () => {
    const results = getFileLines('day-8/input.txt').map(l => l.substr(l.indexOf('|') + 2).split(' '))
    const segments = getFileLines('day-8/input.txt').map(l => l.replace(/ \| /g, ' ').split(' '))

    let sum = 0

    for (let i = 0; i < segments.length; i++) {
        const segment = segments[i]
        const result = results[i]

        const numbers = segmentsToNumbers(segment)

        const digits = result.map(r => numbers.findIndex(n => areSame(n, r)))

        // console.log(result.map(r => r.split("").sort().join("")))
        // console.log(numbers)
        // console.log(digits)
        sum += digits.reduce((total, n, index) => total + (n * Math.pow(10, digits.length - 1 - index)), 0)
    }

    console.log(sum)


}

main()