package y2023d20

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed input.txt
var input string

type module struct {
	label      string
	status     bool
	moduleType rune
	outputs    []string
}

func (m *module) getOutputs() []string {
	return m.outputs
}

func (m *module) addInput(i string) {
}

type flipFlop struct {
	module
}

func (m *flipFlop) getActivation(origin string, newPulse bool) (bool, bool) {
	if !newPulse {
		m.status = !m.status
		return m.status, true
	}

	return m.status, false

}

func (m *flipFlop) getType() rune {
	return '%'
}

func (m *flipFlop) String() string {
	return "flipflop (" + m.label + ") -> " + strings.Join(m.outputs, ", ")
}

type conjunction struct {
	module
	inputs map[string]bool
}

func (m *conjunction) getActivation(origin string, newPulse bool) (bool, bool) {
	//fmt.Println("Conjunction activation:", origin, "->", newPulse)
	m.inputs[origin] = newPulse
	//fmt.Println(m.inputs)

	//  if it remembers high pulses for all inputs,
	//		it sends a low pulse;
	//	otherwise,
	//		it sends a high pulse.
	for _, status := range m.inputs {
		if !status {
			return true, true
		}
	}

	return false, true
}

func (m *conjunction) addInput(i string) {
	m.inputs[i] = false
}

func (m *conjunction) getType() rune {
	return '&'
}

func (m *conjunction) String() string {
	return "conjunction (" + m.label + ") -> " + strings.Join(m.outputs, ", ")
}

type broadcast struct {
	module
}

func (m *broadcast) getActivation(origin string, newPulse bool) (bool, bool) {
	return newPulse, true
}

func (m *broadcast) getType() rune {
	return ' '
}

func (m *broadcast) String() string {
	return "broadcast (" + m.label + ") -> " + strings.Join(m.outputs, ", ")
}

type test struct {
	module
}

func (m *test) getActivation(origin string, newPulse bool) (bool, bool) {
	return false, false
}

func (m *test) getType() rune {
	return 't'
}

func (m *test) String() string {
	return "test (" + m.label + ")"
}

type imodule interface {
	getActivation(origin string, newPulse bool) (bool, bool)
	getType() rune
	getOutputs() []string
	addInput(string)
}

func parseText(text string) (map[string]imodule, string) {
	lines := strings.Split(text, "\n")

	systemM := map[string]imodule{}
	firstModule := ""

	for _, line := range lines {
		originDest := strings.Split(line, " -> ")

		origin := originDest[0]
		outputs := strings.Split(originDest[1], ", ")

		if strings.IndexRune(origin, '%') > -1 {
			label := strings.ReplaceAll(origin, "%", "")

			if firstModule == "" {
				firstModule = label
			}

			systemM[label] = &flipFlop{
				module: module{
					label:   label,
					status:  false,
					outputs: outputs,
				},
			}
			continue
		}

		if strings.IndexRune(origin, '&') > -1 {
			label := strings.ReplaceAll(origin, "&", "")

			if firstModule == "" {
				firstModule = label
			}

			systemM[label] = &conjunction{
				module: module{
					label:   label,
					status:  false,
					outputs: outputs,
				},
				// calculate inputs
				inputs: map[string]bool{},
			}
			continue
		}

		if firstModule == "" {
			firstModule = origin
		}

		systemM[origin] = &broadcast{
			module: module{
				label:   origin,
				status:  false,
				outputs: outputs,
			},
		}

	}

	fmt.Printf("Keys found:\n\t")
	for k, _ := range systemM {
		fmt.Printf("%s\t", k)
	}
	fmt.Println()

	fmt.Println("Adding inputs to conjuctions")
	for k, m := range systemM {

		outputs := m.getOutputs()

		for _, output := range outputs {
			fmt.Println("Evaluating", output)

			_, ok := systemM[output]

			if !ok {
				systemM[output] = &test{
					module{
						label: output,
					},
				}
			}
			// Some outputs are for testing purposes and are not really part of the system
			if _, ok := systemM[output]; ok {
				if systemM[output].getType() == '&' {
					systemM[output].addInput(k)
				}
			}
		}

	}

	return systemM, firstModule

}

type pendingActivation struct {
	origin      string
	destination string
	activation  bool
}

func Run1() {

	fmt.Println(input)
	systemM, first := parseText(input)

	fmt.Printf("%s\n%v\n", first, systemM)

	numberOfClicks := 1000

	totalLowPulses := 0
	totalHighPulses := 0
	for i := 0; i < numberOfClicks; i++ {
		fmt.Println()

		pendingPulses := []pendingActivation{}

		currentActivation := pendingActivation{
			origin:      "button",
			destination: first,
			activation:  false}
		// Click counts as a pulse
		totalLowPulses++

		for {

			module := systemM[currentActivation.destination]
			outputs := module.getOutputs()

			nextPulse, propagate := module.getActivation(currentActivation.origin, currentActivation.activation)

			if propagate {
				for _, output := range outputs {
					// Some outputs are for testing purposes and are not really part of the system
					if _, ok := systemM[output]; ok {

						fmt.Println(currentActivation.destination, "-", nextPulse, "->", output)
						if nextPulse {
							totalHighPulses++
						} else {
							totalLowPulses++
						}
						pendingPulses = append(pendingPulses, pendingActivation{
							destination: output,
							origin:      currentActivation.destination,
							activation:  nextPulse,
						})
					}

				}
			}

			//	fmt.Println("pendingPulses", pendingPulses)
			if len(pendingPulses) == 0 {
				break
			}

			currentActivation = pendingPulses[0]
			pendingPulses = pendingPulses[1:]

		}

	}

	fmt.Println()
	fmt.Println("Total pulses:", totalHighPulses, "*", totalLowPulses, "=", totalHighPulses*totalLowPulses)

}
