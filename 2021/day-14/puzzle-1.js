const { getFileContent } = require('../input-tools')

const [template, tail] = getFileContent('day-14/input.txt', '\r\n\r\n')

const insertionRules = tail.split('\r\n').map(l => l.split(' -> ')).map(l => [l[0], l[1]])

console.log(template)

let currentTemplate = template
let result = ""

for(let times = 0; times < 10; times++) {
    let result = ""
    let chars = currentTemplate.split("")

    for(let i = 0; i< chars.length -1; i++) {
    const [first, next] = chars.slice(i, i + 2)
    // console.log(first, next)

    result += first 

    const param = first + next 
    const repl = insertionRules.find(i => i[0] === param)

    //console.log(repl)

    
    result += repl[1] 
}

    result += chars[chars.length - 1]
    currentTemplate = result
}

//console.log(currentTemplate)

const lettersTime = currentTemplate.split("").reduce((letters, current) => {
    if(!letters[current]) {
        letters[current]= 1
    } else {
        letters[current]++
    }

    return letters
}, {})

console.log(lettersTime)

const letterCount = Object.values(lettersTime)

console.log(Math.max(...letterCount) - Math.min(...letterCount))



