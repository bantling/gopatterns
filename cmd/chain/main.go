package main

import (
	"fmt"
)

// Command defines a command to perform
type Command interface {
	Perform()
}

// PrintCommand prints to the console
type PrintCommand struct {
	msg string
}

// Perform is Command for PrintCommand
func (p PrintCommand) Perform() {
	fmt.Println(p.msg)
}

// BeepCommand beeps
type BeepCommand struct{}

// Perform is Command for BeepCommand
func (p BeepCommand) Perform() {
	fmt.Println("Beep!")
}

// Processor processes one or more Command types, returning true if it processes a particular Command, false if it does not.
type Processor interface {
	Process(Command) bool
}

// AnyProcessor processes any Command
type AnyProcessor struct{}

// Process is Processor for AnyProcessor
func (p AnyProcessor) Process(c Command) bool {
	c.Perform()
	return true
}

// PrintProcessor is Processor for PrintCommand only
type PrintProcessor struct{}

// Process is Processor for PrintProcessor
func (p PrintProcessor) Process(c Command) bool {
	if pc, isa := c.(PrintCommand); isa {
		pc.Perform()
		return true
	}

	return false
}

// Chain tries processing a Command with a series of Processors, and stops at the first Processor that indicates it processed the Command.
type Chain struct {
	processors []Processor
}

// Apply tries each processor until it finds one that applies, if any
func (ch Chain) Apply(cmd Command) {
	fmt.Printf("Attempting to find a processor for Command type %T\n", cmd)

	for _, p := range ch.processors {
		if p.Process(cmd) {
			return
		}
	}

	fmt.Println("Unable to find a processor")
}

func main() {
	var (
		pc = PrintCommand{msg: "Hello, World"}
		bc = BeepCommand{}
		c  = Chain{
			processors: []Processor{
				PrintProcessor{},
			},
		}
	)

	// Will not process BeepCommand, only PrintCommand
	c.Apply(pc)
	c.Apply(bc)

	c.processors = append(c.processors, AnyProcessor{})

	// Will process both commands
	c.Apply(pc)
	c.Apply(bc)
}
