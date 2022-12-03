const { getFileContent } = require('../input-tools')

const [xRange, yRange] = getFileContent('day-17/input.txt')[0].replace('target area: ', '').replace('x=','').replace('y=', '').split(', ').map(i => i.split('..').map(Number))

const start = { x: 0, y: 0}

console.log(start, {xRange, yRange})

const distanceX = [Math.abs(xRange[0]), Math.abs(xRange[1])]

console.log(distanceX)


let maxSteps = 0
let bestInitialSpeedX = 0

const xSpeeds = []
// Brute force. Min. speed in X to reach area. Max of steps 
for(let initialSpeed = 1; initialSpeed < 2000; initialSpeed++) {
    let position = start.x 
    let speed = initialSpeed
    let steps = 0

    while(speed >= 0) {
        position += speed 

        if(position >= Math.min(...distanceX) && position <= Math.max(...distanceX)) {
            xSpeeds.push({initialSpeed: initialSpeed, speed: speed, steps})
        }

        speed--
        steps++
    }
}



const ySpeeds = []

let bestInitialSpeedY = 0
let maxHeight = 0
// Brute force. Min. speed in X to reach area. Max of steps 
for(let initialSpeed = -2000; initialSpeed < 2000; initialSpeed++) {
    let position = start.y 
    let speed = initialSpeed
    let localMaxHeight = 0
    let hasLeftRange = false
    let steps = 0

    while(!hasLeftRange) {
        position += speed
        localMaxHeight = Math.max(localMaxHeight, position)
        
        if(position >= Math.min(...yRange )&& position <= Math.max(...yRange)) {
            ySpeeds.push({initialSpeed: initialSpeed, speed: speed, steps})
        } else {
            const outOfRange = Math.min(...yRange) > position 
            if(outOfRange) {
                hasLeftRange = true
            }
        }
        
        speed--
        steps++

    }

   // console.log(`initial speed:`, initialSpeed, 'finalPosition:', position)
    
            
}
//console.log({maxSteps, bestInitialSpeedX})

//console.log({maxHeight, bestInitialSpeedY})

console.log(xSpeeds)
console.log(ySpeeds)

// Merge both

const speeds = ySpeeds.flatMap(y => xSpeeds.filter(x => (x.speed === 0 && x.steps <= y.steps) || (x.speed > 0 && y.steps === x.steps)).map(x => ([x.initialSpeed, y.initialSpeed]))).filter((s, index, self) => self.findIndex(s2 => s[0] === s2[0] && s[1] === s2[1]) === index)

console.log(speeds.length)
//console.log(speeds)