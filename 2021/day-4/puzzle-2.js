const { getFileLines } = require('../input-tools')

function transpose(arr) {
    let newArr = []
    for (var i = 0; i < arr.length; i++) {
        for (var j = 0; j < arr[i].length; j++) {
            if (!newArr[j]) {
                newArr[j] = []
            }
            newArr[j][i] = arr[i][j]
        }
    }
    return newArr
}

function findLastWiningBoard(numbers, boards) {
    const winingBoards = boards.map(_ => false)
    const masks = boards.map(b => b.map(l => l.map(_ => false)))

    for (let number of numbers) {
        for (let b = 0; b < boards.length; b++) {
            for (let i = 0; i < 5; i++) {
                for (let j = 0; j < 5; j++) {
                    if (boards[b][i][j] === number) {
                        masks[b][i][j] = true
                    }
                }
            }

            let boardHasWon = false
            for (let i = 0; i < 5; i++) {
                boardHasWon = boardHasWon || masks[b].every(l => l[i]) || masks[b][i].every(n => n)
            }

            if (boardHasWon) {
                winingBoards[b] = true
                if (winingBoards.every(b => b)) {
                    return [b, boards[b], masks[b], number]
                }
            }
        }
    }

}

function main() {
    const data = getFileLines("day-4/input.txt")

    const [first, _empty, ...rest] = data


    const numbers = first.split(',').map(Number)
    const boardsCount = rest.length / 6

    const boards = []
    for (var boardIndex = 0; boardIndex < boardsCount; boardIndex++) {
        const start = new Date()
        const boardText = rest.slice(boardIndex * 6, boardIndex * 6 + 6).filter(s => s !== "").map(l => l.trim().replace(/  /g, ' ').split(' ').map(Number))
        boards.push(boardText)
    }


    const [bIndex, wonBoard, boardMasks, lastNumber] = findLastWiningBoard(numbers, boards)
    console.log(numbers)
    console.log(bIndex, wonBoard, lastNumber)

    console.log(boardMasks)
    const totalSum = wonBoard.map((l, lindex) => l.map((r, rindex) => boardMasks[lindex][rindex] ? 0 : r)).reduce((sum, line) => sum + line.reduce((sum2, r) => r + sum2, 0), 0)

    console.log(totalSum, lastNumber)
    console.log(totalSum * lastNumber)






}

main()