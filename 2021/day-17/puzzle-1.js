const { getFileContent } = require('../input-tools')

const [xRange, yRange] = getFileContent('day-17/input.txt')[0].replace('target area: ', '').replace('x=','').replace('y=', '').split(', ').map(i => i.split('..').map(Number))

const start = { x: 0, y: 0}

console.log(start, {xRange, yRange})

const distanceX = [Math.abs(xRange[0]), Math.abs(xRange[1])]

console.log(distanceX)


let maxSteps = 0
let bestInitialSpeedX = 0
// Brute force. Min. speed in X to reach area. Max of steps 
for(let initialSpeed = 1; initialSpeed < Math.max(...distanceX); initialSpeed++) {
    let position = start.x 
    let speed = initialSpeed
    let steps = 0

    while(speed > 0) {
        position += speed 

        if(position >= Math.min(...distanceX)&& position <= Math.max(...distanceX)) {
            if(steps > maxSteps) {
                bestInitialSpeedX = initialSpeed
                maxSteps = steps
            }
        }

        speed--
        steps++
    }
}


let bestInitialSpeedY = 0
let maxHeight = 0
// Brute force. Min. speed in X to reach area. Max of steps 
for(let initialSpeed = 1; initialSpeed < 1000; initialSpeed++) {
    let position = start.y 
    let speed = initialSpeed
    let localMaxHeight = 0
    let hasLeftRange = false

    while(!hasLeftRange) {
        position += speed
        localMaxHeight = Math.max(localMaxHeight, position)
        
        if(position >= Math.min(...yRange )&& position <= Math.max(...yRange)) {
            if(localMaxHeight > maxHeight) {
                bestInitialSpeedY = initialSpeed
                maxHeight = localMaxHeight
            }
        } else {
            const outOfRange = Math.min(...yRange) > position 
            if(outOfRange) {
                hasLeftRange = true
            }
        }
        
        speed--

    }

   // console.log(`initial speed:`, initialSpeed, 'finalPosition:', position)
    
            
}
//console.log({maxSteps, bestInitialSpeedX})

console.log({maxHeight, bestInitialSpeedY})

