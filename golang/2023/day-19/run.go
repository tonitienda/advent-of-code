package y2034d19

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

type part map[rune]int

type operation func(p part) (string, bool)
type operationMaker func(field rune, value int, dest string) func(p part) (string, bool)

type workflows map[string][]operation

func fieldGreaterThan(field rune, value int, dest string) func(p part) (string, bool) {
	return func(p part) (string, bool) {
		if p[field] > value {
			return dest, true
		}

		return "", false
	}
}

func fieldLessThan(field rune, value int, dest string) func(p part) (string, bool) {
	return func(p part) (string, bool) {
		if p[field] < value {
			return dest, true
		}

		return "", false
	}
}

func parseWorkflows(text string) workflows {
	operationsDict := map[rune]operationMaker{
		'>': fieldGreaterThan,
		'<': fieldLessThan,
	}

	lines := strings.Split(text, "\n")

	w := workflows{}

	for _, line := range lines {
		fmt.Println("Parsing:", line)
		s1 := strings.Split(line, "{")

		label := s1[0]
		opsStr := strings.Split(strings.ReplaceAll(s1[1], "}", ""), ",")

		operations := []operation{}

		for _, opStr := range opsStr {

			// The label is the destination
			if !strings.Contains(opStr, ":") {
				operations = append(operations, func(p part) (string, bool) { return opStr, true })
				fmt.Println("\t\t", "=>", opStr)

				continue
			}

			fmt.Println("\tParsing:", opStr)
			field := rune(opStr[0])
			op := rune(opStr[1])

			s2 := strings.Split(opStr[2:], ":")

			valueStr := s2[0]
			dest := s2[1]

			value, err := strconv.Atoi(valueStr)

			if err != nil {
				panic(err)
			}

			fmt.Println("\t\t", string(field), string(op), value, "=>", dest)
			operations = append(operations, operationsDict[op](field, value, dest))
		}

		w[label] = operations
	}

	return w
}

func parseParts(text string) []part {
	parts := []part{}

	for _, line := range strings.Split(text, "\n") {
		cleanStr := strings.ReplaceAll(strings.ReplaceAll(line, "{", ""), "}", "")

		fields := strings.Split(cleanStr, ",")

		p := part{}

		for _, field := range fields {
			x := strings.Split(field, "=")

			label := rune(x[0][0])

			value, err := strconv.Atoi(x[1])

			if err != nil {
				panic(err)
			}

			p[label] = value

		}

		parts = append(parts, p)

	}

	return parts

}

func Run1() {
	blocks := strings.Split(input, "\n\n")
	workflows := parseWorkflows(blocks[0])
	parts := parseParts(blocks[1])

	accepted := []part{}

	for _, part := range parts {
		current := "in"
		acceptedOrRejected := false

		//fmt.Printf("Processing %v", part)

		for !acceptedOrRejected {
			//fmt.Println("Evaluating", current)
			ops := workflows[current]

			if len(ops) == 0 {
				log.Panicf("Wf: %s, has no operations defined", current)
			}

			for _, op := range ops {
				next, ok := op(part)
				//fmt.Println("Next:", next, ok)
				if ok {
					if next == "A" {
						accepted = append(accepted, part)
					}
					acceptedOrRejected = next == "A" || next == "R"
					current = next
					break
				}
			}
		}

	}

	// fmt.Println(workflows)
	// fmt.Println(parts)

	//fmt.Println("Accepted (", len(accepted), ")=", accepted)

	total := 0

	for _, p := range accepted {
		for _, v := range p {
			total += v
		}
	}

	fmt.Println("Total", total)

}

type wfrange struct {
	field rune
	min   int
	max   int
	dest  string
}
type workflowsData map[string][]wfrange

func parseWorkflowsToData(text string) workflowsData {
	operationsDict := map[rune]operationMaker{
		'>': fieldGreaterThan,
		'<': fieldLessThan,
	}

	lines := strings.Split(text, "\n")

	w := workflowsData{}

	for _, line := range lines {
		fmt.Println("Parsing:", line)
		s1 := strings.Split(line, "{")

		label := s1[0]
		opsStr := strings.Split(strings.ReplaceAll(s1[1], "}", ""), ",")

		wfRanges := []wfrange{}

		for _, opStr := range opsStr {

			// The label is the destination
			if !strings.Contains(opStr, ":") {
				wfRanges = append(wfRanges, wfrange{
					field: ' ',
					min:   0,
					max:   4000,
					dest:  opStr,
				})
				fmt.Println("\t\t", "=>", opStr)

				continue
			}

			fmt.Println("\tParsing:", opStr)
			field := rune(opStr[0])
			op := rune(opStr[1])

			s2 := strings.Split(opStr[2:], ":")

			valueStr := s2[0]
			dest := s2[1]

			value, err := strconv.Atoi(valueStr)

			if err != nil {
				panic(err)
			}

			fmt.Println("\t\t", string(field), string(op), value, "=>", dest)
			w[label]
			operations = append(operations, operationsDict[op](field, value, dest))
		}

		w[label] = wfRanges
	}

	return w
}

func Run2() {
	blocks := strings.Split(test, "\n\n")
	workflows := parseWorkflowsToData(blocks[0])

	fmt.Println(workflows)

}
