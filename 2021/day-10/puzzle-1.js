const { getFileLines } = require('../input-tools')


const lines = getFileLines('day-10/input.txt')

const chars = {
    '[': ']',
    '(': ')',
    '{': '}',
    '<': '>'
}

const points = {
    ')': 3,
    ']': 57,
    '}': 1197,
    '>': 25137
}

console.log(lines)
let totalPoints = 0
for (let line of lines) {
    let expected = []
    for (let char of line.split("")) {
        if (chars[char]) {
            expected.push(chars[char])
        } else {
            if (char === expected[expected.length - 1]) {
                expected.pop()
            } else {
                totalPoints += points[char]
                // console.log(line, 'failed', char)
                break
            }
        }
    }
}

console.log(totalPoints)