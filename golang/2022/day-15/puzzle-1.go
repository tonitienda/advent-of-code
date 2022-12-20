package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const TOP = 0
const RIGHT = 1
const BOTTOM = 2
const LEFT = 3

const X = 0
const Y = 0

type Segment struct {
	x1 int
	x2 int
}

type Point struct {
	x int
	y int
}

type Area struct {
	top    Point
	right  Point
	left   Point
	bottom Point
}

func (a Area) String() string {
	return fmt.Sprintf("{ t: [%d, %d], r: [%d, %d], b: [%d, %d], l: [%d, %d] }", a.top.x, a.top.y, a.right.x, a.right.y, a.bottom.x, a.bottom.y, a.left.x, a.left.y)
}

type SensorBeacon struct {
	sensor Point
	beacon Point
}

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func getPoint(str string) Point {
	nums := strings.Split(str, ",")

	x := StrToInt(nums[X])
	y := StrToInt(nums[Y])

	return Point{x: x, y: y}
}
func getSensorAndBeacon(line string) SensorBeacon {
	cleanStr := strings.ReplaceAll(line, "Sensor at x=", "")
	cleanStr = strings.ReplaceAll(cleanStr, " y=", "")
	cleanStr = strings.ReplaceAll(cleanStr, ": closest beacon is at x=", " ")

	parts := strings.Split(cleanStr, " ")

	points := array.Map(parts, getPoint)

	sensor := points[0]
	beacon := points[1]

	return SensorBeacon{sensor, beacon}
}

func overlap(segment1 Segment, segment2 Segment) (Segment, bool) {
	// If overlaps, return the union of both segments
	if segment1.x1 <= segment2.x2 && segment1.x2 >= segment2.x1 {
		min := int(math.Min(float64(segment1.x1), float64(segment2.x1)))
		max := int(math.Max(float64(segment1.x2), float64(segment2.x2)))

		//fmt.Println("Overlaped:", segment1, segment2, "=>", [2]int{min, max})
		return Segment{x1: min, x2: max}, true
	}

	return Segment{}, false
}

func cannotContainABeacon(row int, sensorData []SensorBeacon) int {

	allIntersections := []Segment{}

	for _, sensorBeacon := range sensorData {
		//fmt.Println()
		intersection, intersects := cannotContainBeacon2(row, sensorBeacon)
		// if intersects {
		// 	fmt.Println("intersection", intersection, "intersects", intersects)
		// }
		didOverlap := false

		if intersects {
			// Take into account the overlaps
			for idx, i := range allIntersections {
				if segment, doesOverlap := overlap(i, intersection); doesOverlap {
					allIntersections[idx] = segment
					didOverlap = true
				}
			}
			if !didOverlap {
				allIntersections = append(allIntersections, intersection)
			}
		}
	}

	fmt.Println("allIntersections", allIntersections)
	emptyCells := 0
	for _, intersection := range allIntersections {
		emptyCells += intersection.x2 - intersection.x1
	}

	return emptyCells

}

func cannotContainBeacon2(row int, sensorData SensorBeacon) (Segment, bool) {
	sensor := sensorData.sensor
	beacon := sensorData.beacon

	areaWithoutBeacon, hasArea := getAreaWithoutBeacon(sensor, beacon)

	//fmt.Println(areaWithoutBeacon, hasArea)

	if !hasArea {
		return Segment{}, false
	}

	return intersect(row, areaWithoutBeacon)

}

func intersect(row int, area Area) (Segment, bool) {
	// If row if out of bounds of the area, return empty array
	if area.top.y > row || area.bottom.y < row {
		// Does not interesct
		return Segment{}, false
	}

	centerY := area.left.y

	diff := int(math.Abs(float64(row-centerY)) / 2)
	minX0 := area.left.x
	maxX0 := area.right.x

	minX := minX0 + diff
	maxX := maxX0 - diff

	fmt.Println("area", area, "row", row, "centerY", centerY, "minx0", minX0, "maxX0", maxX0, "diff", diff, "minx", minX, "maxX", maxX)

	return Segment{x1: minX, x2: maxX}, true

}

func getAreaWithoutBeacon(sensor, beacon Point) (Area, bool) {
	distance := getDistance(sensor, beacon)

	// Empty area
	if distance == 0 {
		return Area{}, false
	}

	top := Point{x: sensor.x, y: sensor.y - distance}
	right := Point{x: sensor.x + distance, y: sensor.y}
	bottom := Point{x: sensor.x, y: sensor.y + distance}
	left := Point{x: sensor.x - distance, y: sensor.y}

	return Area{top: top, right: right, bottom: bottom, left: left}, true
}

func getDistance(sensor, beacon Point) int {
	// TODO - Review this
	return int(math.Abs(float64(sensor.x-beacon.x))) + int(math.Abs(float64(sensor.y-beacon.y)))
}

type Input struct {
	file string
	row  int
}

func main() {

	args := Input{file: "test.txt", row: 10}
	//input := Input{file: "input.txt", row: 2000000}

	data := input.GetLines(2022, 15, args.file)
	rowOfInterest := args.row

	sensorData := array.Map(data, getSensorAndBeacon)

	num := cannotContainABeacon(rowOfInterest, sensorData)

	//fmt.Println(sensorData)
	fmt.Println(num)

	// // Tests
	// fmt.Println([4][2]int{}, funct.GetValueB(getAreaWithoutBeacon([2]int{0, 0}, [2]int{0, 0})))
	// fmt.Println([4][2]int{[2]int{0, -1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}, funct.GetValueB(getAreaWithoutBeacon([2]int{0, 0}, [2]int{1, 0})))
	// fmt.Println([4][2]int{[2]int{0, -2}, [2]int{2, 0}, [2]int{0, -2}, [2]int{-2, 0}}, funct.GetValueB(getAreaWithoutBeacon([2]int{0, 0}, [2]int{1, 1})))
}
