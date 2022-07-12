const { getNumbersFromFile } = require('../input-tools')

function main () {

    const numbers = getNumbersFromFile("day-1/input.txt")
    console.log(numbers)
    const incrementsCount = numbers.map((n, index, self) => {
        if(index === 0) {
            return 0
        }

        return self[index] > self[index-1] ? 1 : 0
    }).reduce((acc, n) => acc+n, 0)

    console.log(incrementsCount)


}

main()