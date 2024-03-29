package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

type Robot struct {
	resource string
	costs    map[string]int
	produces int
}

func (robot Robot) prepareRobot(resources map[string]int) (map[string]int, Robot, bool) {
	newResources := resources

	for key, cost := range robot.costs {
		if newResources[key] < cost {
			// Not enough resources
			return resources, Robot{}, false
		}
	}

	// All robots produce 1 unit
	return newResources, Robot{resource: robot.resource, produces: 1}, true
}

func (robot Robot) consume(resources map[string]int) map[string]int {
	newResources := resources

	fmt.Println("\tBuilding ", robot.resource)
	for key, cost := range robot.costs {
		fmt.Println("\tConsuming", key, "=", cost)
		newResources[key] -= cost
	}

	// All robots produce 1 unit
	return newResources
}

// func (robot Robot) collect(resources map[string]int) map[string]int {

// 	resources[robot.resource] += robot.produces

// 	fmt.Printf("1 %s-collecting robot collects %d %s; you now have %d %s.\n", robot.resource, robot.produces, robot.resource, resources[robot.resource], robot.resource)

// 	return resources
// }

type Blueprint struct {
	id     int
	robots map[string]Robot
}

func getRobot(line string) Robot {
	//fmt.Println(line)

	parts := strings.Split(line, " ")

	//fmt.Println(parts)

	cost := StrToInt(parts[1])
	resourceCost := parts[2]

	costs := map[string]int{
		resourceCost: cost,
	}

	if len(parts) == 6 {
		costs[parts[5]] = StrToInt(parts[4])
	}

	return Robot{
		resource: parts[0],
		costs:    costs,
	}
}

func getBluePrint(line string) Blueprint {

	data := strings.Split(strings.ReplaceAll(line, "Blueprint ", ""), ":")

	id := StrToInt(data[0])
	robotLines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(data[1], " Each ", ""), " robot costs", ""), ".")

	robotsBlueprint := map[string]Robot{}
	for _, robotLine := range robotLines {
		if robotLine != "" {
			robot := getRobot(robotLine)
			robotsBlueprint[robot.resource] = robot
		}
	}

	//fmt.Println(robotsBlueprint)

	return Blueprint{
		id:     id,
		robots: robotsBlueprint,
	}
}

func getBuildableDeps(robot Robot, blueprint Blueprint, resources map[string]int) []Robot {

	robots := []Robot{}

	_, _, wasBuilt := robot.prepareRobot(resources)

	if wasBuilt {
		robots = append(robots, robot)
		return robots
	}

	for res, _ := range robot.costs {
		dep := blueprint.robots[res]

		// Robot depends on itself
		if dep.resource == robot.resource {
			continue
		}

		robots = append(robots, getBuildableDeps(dep, blueprint, resources)...)
	}

	return robots

}

func getBuildableRobots(blueprint Blueprint, resources map[string]int) []Robot {

	// Try  to build Geode robots first
	geode := blueprint.robots["geode"]

	robots := getBuildableDeps(geode, blueprint, resources)

	fmt.Println("Buildable robots", robots, "resources", resources)
	return robots
}

func calculateMaxOreOutcome(iterations int, blueprint Blueprint, robots map[string]int, resources map[string]int) int {

	fmt.Println("Resources:", resources)
	if iterations <= 0 {
		return resources["geode"]
	}

	// fmt.Printf("== Minute %d == \n", 25-iterations)
	// fmt.Println("Current geode", resources["geode"])

	// Build robots
	// TODO - We consider that we can build one robot of each kind
	// But in some cases we could build multiple and we need to test all the combinations
	buildableRobots := getBuildableRobots(blueprint, resources)

	// Collect resources
	for key, units := range robots {
		resources[key] += units
	}

	if len(buildableRobots) == 0 {
		// Cloning resources map
		newResources := map[string]int{}
		for key, count := range resources {
			newResources[key] = count
		}

		return calculateMaxOreOutcome(iterations-1, blueprint, robots, newResources)
	}

	// Simulate that we do not build any robot (saving for more expensive ones)

	maxGeodeOpened := calculateMaxOreOutcome(iterations-1, blueprint, robots, resources)
	fmt.Println("WE ALSO REACH THIS CODE")

	//fmt.Println("\nBuildable robots:", len(buildableRobots), "\n")
	for _, robot := range buildableRobots {

		// Cloning resources map
		newResources := map[string]int{}
		for key, count := range resources {
			newResources[key] = count
		}

		robots[robot.resource]++
		//fmt.Printf("The new %s-collecting robot is ready; you now have %d of them.\n", robot.resource, robots[robot.resource])
		geodeOpened := calculateMaxOreOutcome(iterations-1, blueprint, robots, robot.consume(newResources))
		robots[robot.resource]--

		if geodeOpened > maxGeodeOpened {
			maxGeodeOpened = geodeOpened
		}
	}

	return maxGeodeOpened
}

func main() {

	execType := "test"
	bluePrints := array.Map(input.GetLines(2022, 19, execType+".txt"), getBluePrint)

	// Start with first blueprint
	bluePrint := bluePrints[0]
	// We start with one "ore" collecting robot
	robots := map[string]int{"ore": 1}

	initialResources := map[string]int{}

	oreProduced := calculateMaxOreOutcome(12, bluePrint, robots, initialResources)

	// fmt.Println(bluePrints)
	fmt.Println("Geode opened", oreProduced)

}
