const { getFileContent } = require('../input-tools')

const cache = {}

const calculate = (letters, insertionRules, iterations) => {
    //console.log(letters, ' - ', iterations)

    if(cache[letters]) {
       // console.log(cache[letters]) 
        if(cache[letters][iterations]) {
           // console.log(`CACHED~~`)
            return cache[letters][iterations]
        }
    } else {
        cache[letters] = {}
    }

    if (iterations === 0) {
        return {}
    }

    const rule = insertionRules[letters]
    const count1 = calculate(rule[0], insertionRules, iterations - 1)
    const count2 = calculate(rule[1], insertionRules, iterations - 1)

    const count = mergeCounts(count1, count2)


    const letterToIncrement = rule[2]
    if (!count[letterToIncrement]) {
        count[letterToIncrement] = 1
    } else {
        count[letterToIncrement]++
    }

    
    cache[letters][iterations] = count

    return count


}

const mergeCounts = (count1, count2) => {

    const count = { ...count1 }

    for (let key of Object.keys(count2)) {

        if (!count[key]) {
            count[key] = count2[key]
        } else {
            count[key] += count2[key]
        }
    }

    return count
}

function execute(template, insertionRules) {
    // Initial count is the number of letters in the template


    let count = template.split("").reduce((counts, char) => {
        if(!counts[char]){
            counts[char] = 1
        } else {
            counts[char]++
        }
        return counts
    }, {})
    let letters = template.split("")

    for(let i=0; i<letters.length-1; i++) {
        const start = new Date()
        count = mergeCounts(count, calculate(letters[i]+letters[i+1], insertionRules, 40))
        console.log(i, new Date() - start, 'ms')
    }

    console.log(count)
    
const letterCount = Object.values(count)

console.log(Math.max(...letterCount) - Math.min(...letterCount))
}


const [template, tail] = getFileContent('day-14/input.txt', '\r\n\r\n')

const insertionRules = tail.split('\r\n').map(l => l.split(' -> ')).map(l => [l[0], [l[0].split("")[0] + l[1], l[1] + l[0].split("")[1], l[1]]]).reduce((rules, rule) => ({...rules, [rule[0]]: rule[1]}), {})
//console.log(insertionRules)
const start = new Date()
execute(template, insertionRules)
console.log(new Date() - start, 'ms')
