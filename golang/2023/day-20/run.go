package y2023d20

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test1.txt
var test1 string

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
	m.inputs[origin] = newPulse

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
				fmt.Printf("Output not found: '%s'", output)
			}
			if systemM[output].getType() == '&' {
				systemM[output].addInput(k)
			}
		}

	}

	return systemM, firstModule

}

func Run1() {

	fmt.Println(test1)

	systemM, first := parseText(test1)

	fmt.Printf("%s\n%v\n", first, systemM)

	numberOfClicks := 1

	currentModule := first
	totalPulses := 0
	for i := 0; i < numberOfClicks; i++ {
		pendingPulses := []struct {
			string
			bool
		}{}

		module := systemM[currentModule]

		outputs := module.getOutputs()

		nextPulse, propagate := module.getActivation("", false)

		if propagate {
			for _, output := range outputs {
				fmt.Println(currentModule, "-", nextPulse, "->", output)
				totalPulses++
				pendingPulses = append(pendingPulses, struct {
					string
					bool
				}{
					string: output,
					bool:   nextPulse,
				})

			}
		}

		fmt.Println("pendingPulses", pendingPulses)

	}

}
