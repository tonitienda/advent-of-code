const { getFileLines } = require('../input-tools')


const lines = getFileLines('day-10/input.txt')

const chars = {
    '[': ']',
    '(': ')',
    '{': '}',
    '<': '>'
}

const points = {
    ')': 1,
    ']': 2,
    '}': 3,
    '>': 4
}

// console.log(lines)
let totalPoints = []
for (let line of lines) {
    let expected = []
    for (let char of line.split("")) {
        if (chars[char]) {
            expected.push(chars[char])
        } else {
            if (char === expected[expected.length - 1]) {
                expected.pop()
            } else {
                // totalPoints += points[char]
                // console.log(line, 'failed', char)
                expected = []
                break
            }
        }
    }
    if (expected.length > 0) {
        totalPoints.push({ line: expected.reverse().join(""), points: expected.reduce((total, char) => (total * 5) + points[char], 0)})
    }
}

console.log(totalPoints.sort((a,b) => a.points - b.points)[Math.floor(totalPoints.length / 2)])