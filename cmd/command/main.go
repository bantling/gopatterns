package main

import (
	"fmt"
)

// receiver performs useful function
// command contains receiver and invokes receiver with hard-coded params
// invoker executes command, optional bookkeeping
// client has commands, receivers, and invokers.
//   assigns receivers to commands with params to use.
//   decides which commands to execute as needed.
//   executes commands by passing them to invoker.

// SmartBulb is a receiver that can be on or off, and select one of 3 colours
type SmartBulb struct {
	ID     int
	On     bool
	Colour string
}

// NewSmartBulb constructs a SmartBulb
func NewSmartBulb(id int) *SmartBulb {
	return &SmartBulb{ID: id}
}

// Switch switches bulb on or off
func (s *SmartBulb) Switch(on bool) {
	s.On = on
	fmt.Println("Bulb", s.ID, "switched", on)
}

// Changes sets colour
func (s *SmartBulb) Change(colour string) {
	s.Colour = colour
	fmt.Println("Bulb", s.ID, "coloured", colour)
}

// Command is an interface for a method of no args
type Command interface {
	Execute()
}

// CommandFunc is an adapter to allow use of ordinary functions as commands
type CommandFunc func()

// Execute calls c()
func (c CommandFunc) Execute() {
	c()
}

// SwitchOn generates a Command for a bulb to turn it on
func SwitchOn(s *SmartBulb) Command {
	return CommandFunc(func() { s.Switch(true) })
}

// SwitchOff generates a Command for a bulb to turn it off
func SwitchOff(s *SmartBulb) Command {
	return CommandFunc(func() { s.Switch(false) })
}

// ColourRed generates a Command for a bulb to set the colour red
func ColourRed(s *SmartBulb) Command {
	return CommandFunc(func() { s.Change("Red") })
}

// ColourGreen generates a Command for a bulb to set the colour green
func ColourGreen(s *SmartBulb) Command {
	return CommandFunc(func() { s.Change("Green") })
}

// ColourBlue generates a Command for a bulb to set the colour blue
func ColourBlue(s *SmartBulb) Command {
	return CommandFunc(func() { s.Change("Blue") })
}

// SequenceInvoker is the invoker, it executes a sequence of commands and tracks how many times it has been called
type SequenceInvoker struct {
	commands []Command
	count    int
}

// NewSequenceInvoker constructs a SequenceInvoker
func NewSequenceInvoker() *SequenceInvoker {
	return &SequenceInvoker{}
}

// WithCommand adds a command to the sequence
func (s *SequenceInvoker) WithCommand(cmd Command) *SequenceInvoker {
	s.commands = append(s.commands, cmd)
	return s
}

// InvokeAndCount executes the sequence of commands and increments the count if no command panics
func (ci *SequenceInvoker) InvokeAndCount() {
	for _, cmd := range ci.commands {
		cmd.Execute()
	}
	ci.count++
	fmt.Println("Sequence invoked", ci.count, "times")
}

// Client is the client
type Client struct {
	invokers map[int]*SequenceInvoker
}

// NewClient constructs a Client
func NewClient() *Client {
	return &Client{invokers: map[int]*SequenceInvoker{}}
}

// WithSequence adds a sequence to the client
func (c *Client) WithSequence(seqNo int, seq *SequenceInvoker) *Client {
	c.invokers[seqNo] = seq
	return c
}

func (c Client) PerformSequence(seqNo int) {
	c.invokers[seqNo].InvokeAndCount()
	fmt.Println("Performed sequence", seqNo)
}

func main() {

}
