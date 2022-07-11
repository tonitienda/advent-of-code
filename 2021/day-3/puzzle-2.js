const { getFileLines } = require('../input-tools')

function transpose(arr) {
    let newArr = []
    for(var i = 0; i< arr.length; i++) {
        for(var j = 0; j < arr[i].length; j++){
            if(!newArr[j]) {
                newArr[j] = []
            }              
            newArr[j][i] = arr[i][j]
        }
    }
    return newArr
}

function main () {
    const data = getFileLines("day-3/input.txt").map(l => l.split('').map(Number))

    let mostCommonData = data.map(r => r)
    let index = 0
    while(mostCommonData.length > 1) {
        const calculated = transpose(mostCommonData).map(row =>row.reduce((sum, i) => sum+i,0) >= (row.length /2))

        const mostCommon = calculated.map(c => c ? 1 : 0)

        mostCommonData = mostCommonData.filter(r => r[index] === mostCommon[index])
        index++
    }

    let lessCommonData = data.map(r => r)
    index = 0
    while(lessCommonData.length > 1) {
        const calculated = transpose(lessCommonData).map(row =>row.reduce((sum, i) => sum+i,0) >= (row.length /2))
        const lessCommon = calculated.map(c => c ? 0 : 1)

        lessCommonData = lessCommonData.filter(r => r[index] === lessCommon[index])
        index++
    }

    const  oxygenGenerator= parseInt(mostCommonData[0].reduce((acc, n) => acc + n, ""),2);
    const  co2Scrubber= parseInt(lessCommonData[0].reduce((acc, n) => acc + n, ""),2);


    console.log(oxygenGenerator)
    console.log(co2Scrubber)

    console.log(oxygenGenerator * co2Scrubber)
}

main()