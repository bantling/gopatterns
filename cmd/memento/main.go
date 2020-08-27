package main

import (
	"fmt"
)

// Originator contains the set to edit and undo
type Originator struct {
	data map[string]bool
}

// Memento is the memento object that contains the info to return the set to a previous state
type Memento struct {
	data map[string]bool
}

// NewOriginator constructs an Originator
func NewOriginator() *Originator {
	return &Originator{data: map[string]bool{}}
}

// Add adds items to the set
func (o *Originator) Add(items []string) {
	for _, item := range items {
		o.data[item] = true
	}
}

// Remove removes items from the set
func (o *Originator) Remove(items []string) {
	for _, item := range items {
		delete(o.data, item)
	}
}

// Memento returns a memento object that can return the set to the original state
func (o Originator) Memento() Memento {
	dataCopy := map[string]bool{}
	for k, _ := range o.data {
		dataCopy[k] = true
	}

	return Memento{data: dataCopy}
}

// Apply a memento to restore set to a previous state
func (o *Originator) SetMemento(m Memento) {
	o.data = map[string]bool{}
	for k, _ := range m.data {
		o.data[k] = true
	}
}

// Caretaker
type Caretaker struct {
	o        *Originator
	mementos []Memento
}

// NewCaretaker constructs a Caretaker
func NewCaretaker() *Caretaker {
	return &Caretaker{o: NewOriginator()}
}

// KeepTrackOf keeps track of some items for later retrieval in random order.
func (c *Caretaker) KeepTrackOf(items ...string) {
	// First get a memento of current state, then make changes
	c.mementos = append(c.mementos, c.o.Memento())
	c.o.Add(items)
}

// LoseTrackOf loses track of some items, so they are no longer available for retrieval later.
// There does not have to be any relationship between the items passed to KeepTrackOf and LoseTrackOf.
// The items passed to LoseTrackOf do not have to exist in the current set.
func (c *Caretaker) LoseTrackOf(items ...string) {
	// First get a memento of current state, then make changes
	c.mementos = append(c.mementos, c.o.Memento())
	c.o.Remove(items)
}

// Items returns the current set of items in random order
func (c Caretaker) Items() []string {
	var items []string
	for k, _ := range c.o.data {
		items = append(items, k)
	}

	return items
}

// Undo undoes a call to KeepTrackOf/LoseTrackOf.
// The effect is as if each call to KeepTrackOf or LoseTrackOf gets added to the top of a stack, and each call to Undo removes and reverses the operation performed by the top item on the stack.
// Panics if there are no operations remaining to undo (which implies the set of items is empty).
func (c *Caretaker) Undo() {
	if len(c.mementos) == 0 {
		panic(fmt.Errorf("No more item sets remain to be removed"))
	}

	lastIdx := len(c.mementos) - 1
	lastMemento := c.mementos[lastIdx]
	c.o.SetMemento(lastMemento)
	c.mementos = c.mementos[:lastIdx]
}

func main() {
	careTaker := NewCaretaker()
	fmt.Println("empty", careTaker.Items())

	careTaker.KeepTrackOf("1", "2", "3")
	fmt.Println("added 1,2,3", careTaker.Items())

	careTaker.LoseTrackOf("2")
	fmt.Println("lost 2", careTaker.Items())

	careTaker.Undo()
	fmt.Println("undo losing 2", careTaker.Items())

	careTaker.Undo()
	fmt.Println("undo adding 1,2,3", careTaker.Items())

	// Die
	careTaker.Undo()
}
