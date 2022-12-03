const { getNumbersFromFile } = require('../input-tools')

const main = () => {
const fish = getNumbersFromFile('day-6/input.txt', ',')

let fishCount = [0,0,0,0,0,0,0,0,0].map((_, idx) => fish.filter(n => n === idx).length)

 for(let i=0; i< 256; i++) {
     const fishToAdd = fishCount[0]

     fishCount = fishCount.map((_, idx, self) => self[idx+1])
     fishCount[6] += fishToAdd
     fishCount[8] = fishToAdd
 }

console.log(fishCount.reduce((sum, n) => sum + n,0))
}

main()  