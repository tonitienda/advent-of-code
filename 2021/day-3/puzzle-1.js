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

    const calculated = transpose(data).map(row =>row.reduce((sum, i) => sum+i,0) > (row.length /2))

    console.log(calculated)

    const gamma = parseInt(calculated.map(c => c ? 1 : 0).reduce((acc, n) => acc + n, ""),2);
    const epsilon = parseInt(calculated.map(c => c ? 0 : 1).reduce((acc, n) => acc + n, ""),2);

    
    console.log(gamma)
    console.log(epsilon)

    console.log(gamma * epsilon)
}

main()