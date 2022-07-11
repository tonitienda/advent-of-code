const { getNumbersFromFile } = require('../input-tools')

const main = () => {
const fish= getNumbersFromFile('day-6/input.txt', ',')

for(let i=0; i< 80; i++) {
    let fishToAdd = 0

    for(let j=0; j<fish.length;j++) {
        if(fish[j] ===0) {
            fishToAdd++
            fish[j]=6
        } else {
            fish[j]--
        }
    }

    for(let k=0; k<fishToAdd; k++) {
        fish.push(8)
    }
}

console.log(fish.length)
}

main()  