const { getNumbersFromFile } = require('../input-tools')

function main () {

    const numbers = getNumbersFromFile("day-1/input.txt")
    const windowedIncrements = numbers.map((n, index, self) => {
        if(index > self.length - 3) {
            return -1
        }

        return self[index] + self[index + 1]  + self[index + 2]
    }).filter(n => n !== -1)
    
    const incrementsCount = windowedIncrements.map((n, index, self) => {
        if(index === 0) {
            return 0
        }

        return self[index] > self[index-1] ? 1 : 0
    }).reduce((acc, n) => acc+n, 0)

    console.log(incrementsCount)



}

main()