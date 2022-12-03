const { getFileLines } = require('../input-tools')

const split = (list) => {

    let result = list

    let itemToSplit = result.findIndex(n => n.value >= 10)
    let wasSplit = false

    if (itemToSplit > -1) {

        // console.log('Split:')
        // printList(result)
        // console.log(`${Array(itemToSplit).fill("   ").join("")}^^^^`)

        wasSplit = true
        let item = result[itemToSplit]

        const leftList = result.slice(0, Math.max(itemToSplit, 0))
        const rightList = result.slice(itemToSplit + 1)

        result = [
            ...leftList,
            { value: Math.floor(item.value / 2), depth: item.depth + 1 },
            { value: Math.ceil(item.value / 2), depth: item.depth + 1 },
            ...rightList
        ]


        //console.log('Split:')
        //printList(result)

        itemToSplit = result.findIndex(n => n.value >= 10)


    }

    return [wasSplit, result]

}
const explode = (list) => {
    // Find the left most pair with depth 4
    let maxDepth = Math.max(...list.map(n => n.depth))
    if (maxDepth < 4) {
        return [false, list]
    }

    let pairtToExplode = list.findIndex((n, index, self) => n.depth >= maxDepth && self[index + 1] && n.depth === self[index + 1].depth)

    let result = list
    let wasExploded = false

    while (pairtToExplode > -1) {

        // console.log('Explode:')
        // printList(result)
        // console.log(`${Array(pairtToExplode).fill("   ").join("")}^^^^`)
        wasExploded = true
        const leftNode = result[pairtToExplode]
        const rightNode = result[pairtToExplode + 1]

        const nodeToAddLeft = result[pairtToExplode - 1]
        const nodeToAddRight = result[pairtToExplode + 2]

        const leftList = result.slice(0, Math.max(pairtToExplode - 1, 0))
        const rightList = result.slice(pairtToExplode + 3)

        const explodedPart = [
            nodeToAddLeft ? { ...nodeToAddLeft, value: nodeToAddLeft.value + leftNode.value } : null,
            { value: 0, depth: leftNode.depth - 1 },
            nodeToAddRight ? { ...nodeToAddRight, value: nodeToAddRight.value + rightNode.value } : null,
        ].filter(n => n !== null)


        result = [
            ...leftList,
            ...explodedPart,
            ...rightList
        ]


        // printList(result)

        pairtToExplode = result.findIndex(n => n.depth === 4)

    }

    return [wasExploded, result]
}


const makeList = (expression, depth = 0) => {
    const [l, r] = expression

    return [
        ...(Array.isArray(l) ? makeList(l, depth + 1) : [{ value: l, depth }]),
        ...(Array.isArray(r) ? makeList(r, depth + 1) : [{ value: r, depth }])
    ]
}

const magnitude = (list) => {
    printList(list)

    let maxDepth = Math.max(...list.map(n => n.depth))
    let pairtToExplode = list.findIndex((n, index, self) => n.depth >= maxDepth && self[index + 1] && n.depth === self[index + 1].depth)
console.log({pairtToExplode})
    const leftList = list.slice(0, Math.max(pairtToExplode - 1, 0))
    const rightList = list.slice(pairtToExplode + 3)


    const result = [
        ...leftList,
        { value: list[pairtToExplode] * 3 + list[pairtToExplode + 1] * 2 },
        ...rightList
    ]

    printList(result)
    //return result
}

const printList = (list) => {
    list.forEach(i => process.stdout.write(`${i.value}, `))
    console.log()
}

const processList = (list) => {
    let currentList = list
    let wasSplit = false
    let wasExploded = false
    do {
        let explRes = explode(currentList)
        let splRes = split(explRes[1])

        wasSplit = splRes[0]
        wasExploded = explRes[0]

        currentList = splRes[1]
    } while (wasSplit || wasExploded)

    return currentList
}

const homeworkLines = getFileLines('day-18/test.txt')
let [currentList, ...rest] = homeworkLines.map(l => makeList(eval(l)))


while (rest.length > 0) {
    let [next, ...tail] = rest
    rest = tail

    //console.log(currentList)
    currentList = currentList.concat(next).map(i => ({ ...i, depth: i.depth + 1 }))

    console.log('-------')
    printList(currentList)
    currentList = processList(currentList)
    console.log('-----')

}


printList(currentList)

console.log(magnitude(currentList))
