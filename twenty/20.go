package twenty

import (
	"aoc"
	"strings"
)

type signal int

const (
	nothing   signal = 0
	lowPulse  signal = 1
	highPulse signal = 2
)

type emission struct {
	source string
	sig    signal
}

type module interface {
	receive(c *circuit, from string, pulse signal) signal
	outConnections() []string
}

type circuit struct {
	modules     map[string]module
	connections map[string][]string
}

type flipFlop struct {
	name        string
	on          bool
	connections []string
}

func (f *flipFlop) outConnections() []string {
	return f.connections
}

func (f *flipFlop) receive(_ *circuit, _ string, pulse signal) signal {
	if pulse == highPulse {
		return nothing
	}
	if f.on {
		f.on = false
		return lowPulse
	} else {
		f.on = true
		return highPulse
	}
}

type conjunction struct {
	name        string
	memory      map[string]signal
	connections []string
}

func (m *conjunction) outConnections() []string {
	return m.connections
}

func (m *conjunction) receive(c *circuit, from string, pulse signal) signal {
	inputs := c.connections[m.name]
	m.memory[from] = pulse
	if len(m.memory) != len(inputs) {
		return highPulse
	}

	for _, s := range m.memory {
		if s != highPulse {
			return highPulse
		}
	}

	// all are high so send low
	return lowPulse
}

type broadcaster struct {
	name        string
	connections []string
}

func (b *broadcaster) outConnections() []string {
	return b.connections
}

func (b *broadcaster) receive(_ *circuit, _ string, s signal) signal {
	return s
}

func send(c *circuit, from string, s signal, f module) []emission {
	var result []emission
	for _, conn := range f.outConnections() {
		nextSignal := c.modules[conn].receive(c, from, s)
		result = append(result, emission{source: conn, sig: nextSignal})
	}
	return result
}

func calcPulses(file string, presses int) int {
	cir := parse(file)
	c := &cir

	totalHigh, totalLow := 0, 0
	for i := 0; i < presses; i++ {
		//println("--------press---------")
		emissions := []emission{{source: "broadcaster", sig: lowPulse}}
		totalLow++
		for len(emissions) > 0 {
			current := emissions[0]
			emissions = emissions[1:]
			currentModule := c.modules[current.source]
			for _, con := range currentModule.outConnections() {
				if current.sig == lowPulse {
					//println(current.source, "low->", con)
					totalLow++
				}
				if current.sig == highPulse {
					//println(current.source, "high->", con)
					totalHigh++
				}
				connectedModule, p := c.modules[con]
				if !p {
					continue
				}
				newSig := connectedModule.receive(c, current.source, current.sig)
				if newSig != nothing {
					emissions = append(emissions, emission{source: con, sig: newSig})
				}

			}
		}
	}
	return totalLow * totalHigh
}

func findRxLowSend(file string) int {
	cir := parse(file)
	c := &cir

	totalHigh, totalLow := 0, 0
	found := false
	for !found {
		//println("--------press---------")
		emissions := []emission{{source: "broadcaster", sig: lowPulse}}
		totalLow++
		for len(emissions) > 0 {
			current := emissions[0]
			emissions = emissions[1:]
			currentModule := c.modules[current.source]
			for _, con := range currentModule.outConnections() {
				if current.sig == lowPulse {
					//println(current.source, "low->", con)
					totalLow++
					if con == "rx" {
						found = true
					}
				}
				if current.sig == highPulse {
					//println(current.source, "high->", con)
					totalHigh++
				}
				connectedModule, p := c.modules[con]
				if !p {
					continue
				}
				newSig := connectedModule.receive(c, current.source, current.sig)
				if newSig != nothing {
					emissions = append(emissions, emission{source: con, sig: newSig})
				}

			}
		}
	}
	return totalLow * totalHigh
}

func parse(file string) circuit {
	scanner := aoc.OpenScanner(file)
	nodes := make(map[string]module)
	for scanner.Scan() {
		line := scanner.Text()
		iAndO := strings.Split(line, "->")
		namePart := strings.TrimSpace(iAndO[0])
		firstChar := namePart[0]
		output := buildOutput(iAndO[1])
		name := namePart[1:]
		if firstChar == '%' {
			nodes[name] = &flipFlop{name: name, on: false, connections: output}
		} else if firstChar == '&' {
			nodes[name] = &conjunction{name: name, memory: make(map[string]signal), connections: output}
		} else {
			nodes[namePart] = &broadcaster{name: namePart, connections: output}
		}
	}

	return circuit{modules: nodes, connections: buildConnections(nodes)}
}

func buildOutput(output string) []string {
	split := strings.Split(output, ",")
	var result = make([]string, len(split))
	for i, s := range split {
		result[i] = strings.TrimSpace(s)
	}
	return result
}

func buildConnections(modules map[string]module) map[string][]string {
	var result = make(map[string][]string)
	for n, m := range modules {
		for _, c := range m.outConnections() {
			result[c] = append(result[c], n)
		}
	}
	return result
}
