package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func getNumbersInLine(line string) []int {
	nums := []int{}

	for _, numstr := range strings.Split(line, " ") {
		v, err := strconv.Atoi(numstr)

		if err != nil {
			panic(err)
		}
		nums = append(nums, v)
	}

	return nums
}

func getSeeds(line string) []int {
	seedsStr := strings.ReplaceAll(line, "seeds: ", "")
	return getNumbersInLine(seedsStr)

}

func getMapFromBlock(block string) [][3]int {

	refMap := [][3]int{}

	lines := strings.Split(block, "\n")

	// We can ignore the first line because it contains the title of the block
	for _, line := range lines[1:] {
		nums := getNumbersInLine(line)

		refMap = append(refMap, [3]int{nums[0], nums[1], nums[2]})
	}

	return refMap
}

func getCorrespondence2(indirections *[][3]int, s int) int {
	for _, indirection := range *indirections {
		// The value is within range
		if s >= indirection[1] && s < indirection[1]+indirection[2] {
			return indirection[0] + (s - indirection[1])
		}
	}

	return s
}

var humToLocCache map[int]int
var seedToSoilCache map[int]int
var tempToHumCache map[int]int
var lightToTempCache map[int]int
var waterToLightCache map[int]int
var fertToWaterCache map[int]int
var soilToFertCache map[int]int

func init() {
	humToLocCache = map[int]int{}
	seedToSoilCache = map[int]int{}
	tempToHumCache = map[int]int{}
	lightToTempCache = map[int]int{}
	waterToLightCache = map[int]int{}
	fertToWaterCache = map[int]int{}
	soilToFertCache = map[int]int{}
}

// func getCachedCorrespondence(cache *map[int]int, m [][3]int) func(s int) int {
// 	return func(s int) int {
// 		val, ok := (*cache)[s]

// 		if ok {
// 			fmt.Println("Cache hit!")
// 			return val
// 		}

// 		res := getCorrespondence(m, s)

// 		(*cache)[s] = res

// 		return res
// 	}

// }

func main() {

	blocks := strings.Split(input, "\n\n")

	//fmt.Println("blocks", blocks)

	// Assuming that the blocks are always in the same order
	seeds := getSeeds(blocks[0])

	seedsToSoil := getMapFromBlock(blocks[1])
	soilToFertilizer := getMapFromBlock(blocks[2])
	fertilizerToWater := getMapFromBlock(blocks[3])
	waterToLight := getMapFromBlock(blocks[4])
	lightToTemperature := getMapFromBlock(blocks[5])
	temperatureToHumidity := getMapFromBlock(blocks[6])
	humidityToLocation := getMapFromBlock(blocks[7])

	// fmt.Println("Seeds:", seeds)
	// fmt.Println("Seeds to Soil:", seedsToSoil)
	// fmt.Println("Seeds to Soil:", soilToFertilizer)
	// fmt.Println("Seeds to Soil:", fertilizerToWater)
	// fmt.Println("Seeds to Soil:", waterToLight)
	// fmt.Println("Seeds to Soil:", lightToTemperature)
	// fmt.Println("Seeds to Soil:", temperatureToHumidity)
	// fmt.Println("Seeds to Soil:", humidityToLocation)

	// getHumidityToLocation := getCachedCorrespondence(&humToLocCache, humidityToLocation)
	// getSeedtoSoil := getCachedCorrespondence(&seedToSoilCache, seedsToSoil)
	// getTemperatureToHumidity := getCachedCorrespondence(&tempToHumCache, temperatureToHumidity)
	// getLightToTemperature := getCachedCorrespondence(&lightToTempCache, lightToTemperature)
	// getWaterToLight := getCachedCorrespondence(&waterToLightCache, waterToLight)
	// getFertilizerToWater := getCachedCorrespondence(&fertToWaterCache, fertilizerToWater)
	// getSoilToFert := getCachedCorrespondence(&soilToFertCache, soilToFertilizer)

	lowestLocation := math.MaxInt
	for i := 0; i < len(seeds); i = i + 2 {
		fmt.Println("Cheking seeds:", seeds[i], "-", seeds[i]+seeds[i+1])
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			//fmt.Println("Seed:", seed)

			//fmt.Println("seed", seed, "soil", getCorrespondence(seedsToSoil, seed))

			// location := getHumidityToLocation(
			// 	getTemperatureToHumidity(
			// 		getLightToTemperature(
			// 			getWaterToLight(
			// 				getFertilizerToWater(
			// 					getSoilToFert(
			// 						getSeedtoSoil(seed)))))))

			location := getCorrespondence2(&humidityToLocation,
				getCorrespondence2(&temperatureToHumidity,
					getCorrespondence2(&lightToTemperature,
						getCorrespondence2(&waterToLight,
							getCorrespondence2(&fertilizerToWater,
								getCorrespondence2(&soilToFertilizer,
									getCorrespondence2(&seedsToSoil, seed)))))))

			//fmt.Println("location", location)
			lowestLocation = int(math.Min(float64(lowestLocation), float64(location)))
		}

	}
	fmt.Println(lowestLocation)

}
