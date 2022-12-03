const { getFileContent } = require('../input-tools')


const HexMap = {
    "0": "0000",
    "1": "0001",
    "2": "0010",
    "3": "0011",
    "4": "0100",
    "5": "0101",
    "6": "0110",
    "7": "0111",
    "8": "1000",
    "9": "1001",
    "A": "1010",
    "B": "1011",
    "C": "1100",
    "D": "1101",
    "E": "1110",
    "F": "1111",
}

const toBinary = (hex) => {
    return hex.split("").map(c => HexMap[c]).join("")
}

const getFirst = (str, length) => {
    return [str.substring(0, length), str.substring(length)]
}

const processHeader = (binary) => {
    const [version, rest1] = getFirst(binary, 3)
    const [type, rest2] = getFirst(rest1, 3)


    return [{
        version: parseInt(version, 2), type: parseInt(type, 2)
    }, rest2]
}

const processLiteralBody = (binary) => {
    //console.log(`Literal packet:`, binary)
    let i = 0
    let number = ""
    let isLast = false
    const blockLength = 5
    while (!isLast) {
        let block = binary.substring(i * blockLength, blockLength + i * blockLength)

        number += block.substring(1, blockLength)
        isLast = block[0] === "0"
        i++
    }

    //console.log(`Literal package:`, number)
    // console.log("Remove trailing zeroes:",  (i * blockLength), '=>',  (i * blockLength + (4 - (i * blockLength) % 4) %4))

    // const blockToIgnore =  (i * blockLength + (4 - (i * blockLength) % 4) %4)
    const blockToIgnore = i * blockLength


    //console.log(`Literal packet: value:`,parseInt(number, 2), 'rest:', binary.substring(blockToIgnore), 'removed:', blockToIgnore)
    return [parseInt(number, 2), binary.substring(blockToIgnore)]
}

const processOperator1Body = (binary) => {
    //console.log(`processOperator1Body:`, binary)

    const [packetsNumber, rest] = getFirst(binary, 11)


    // console.log(`processOperator1Body: packetsNumber:`, packetsNumber, '(', parseInt(packetsNumber, 2), ')')

    return processPacket(rest, parseInt(packetsNumber, 2))
}

const processOperator0Body = (binary) => {
    // console.log(`processOperator0Body:`, binary)
    const [bitLength, rest] = getFirst(binary, 15)

    const length = parseInt(bitLength, 2)

    //console.log(`processOperator0Body: bitLength:`, bitLength, '(', length, ')')
    const [packets, rest2] = getFirst(rest, length)

    //console.log(`processOperator0Body:`, rest2)

    const [value, _] = processPacket(packets)

    return [value, rest2]
}

const processOperatorBody = (binary) => {
    const [type, rest] = getFirst(binary, 1)

    if (type === "1") {
        return processOperator1Body(rest)
    } else {
        return processOperator0Body(rest)
    }
}

const calculateValue = (packets) => {
    return packets.reduce((sum, p) => sum + getPacketValue(p), 0)
}


    // 0 sum
    // 1 product
    // 2 minimum
    // 3 maximum
    // 5 greater ? 1 : 0  (2 items)
    // 6 less than ? 1 : 0 (2 items
    // 7 equals ? 1 : 0 (2 items)
const Operations = {
    0: (values) => values.reduce((total, v) => total + v, 0),
    1: (values) => values.reduce((total, v) => total * v, 1),
    2: (values) => Math.min(...values),
    3: (values) => Math.max(...values),
    5: ([v1, v2]) => v1 > v2 ? 1 : 0,
    6: ([v1, v2]) => v1 < v2 ? 1 : 0,
    7: ([v1, v2]) => v1 === v2 ? 1 : 0
}

const getPacketValue = (packet) => {
    if (packet.type === 4) {
        return packet.value || 0
    }

    const values = packet.subpackets.map(getPacketValue)
if(isNaN(packet.type)) {
    return 0
}
console.log(packet)
    return Operations[packet.type](values)
}


const processPacket = (binary, packetsNumber) => {
    let rest = binary

    const packets = []
    //console.log(`out - packetsNumber:`,packetsNumber, `packets:`, packets.length)

    while (rest.length > 0 && (!packetsNumber || packets.length < packetsNumber)) {
        //   console.log(`packetsNumber:`,packetsNumber, `packets:`, packets.length)
        //console.log(rest)
        const [header, rest2] = processHeader(rest)
        rest = rest2

        //console.log(header)
        if (header.type === 4) {
            const [value, rest3] = processLiteralBody(rest)
            packets.push(
                {
                    ...header,
                    value,
                }
            )
            rest = rest3
            //console.log(rest.length)
        } else {
            // console.log(`processing operator`)
            const [subpackets, rest4] = processOperatorBody(rest)

            packets.push(
                {
                    ...header,
                    subpackets: subpackets.filter(s => s !== null)
                }
            )
            rest = rest4

        }
    }
    //console.log('exit')

    return [packets, rest]
}

const printPacket = ({ type, version, value, subpackets }, level = 0) => {
    const prefix = new Array(level).fill('\t').join("")
    const str = `${prefix}version=${version} type=${type} value=${value}`

    //console.log(str)
    if (subpackets) {
        subpackets.forEach(s => printPacket(s, level + 1))
    }
}


// TEST1 = 16
// TEST2 = 12
// TEST3 = 23
// TEST4 = 31
const packet = getFileContent('day-16/input.txt')[0]
const binary = toBinary(packet)


// console.log(packet)
// console.log(binary)

const [packets, rest] = processPacket(binary)


//const [packets, rest] = processPacket("11101110000000001101010000001100100000100011000001100000")
packets.forEach(printPacket)

console.log(calculateValue(packets))
