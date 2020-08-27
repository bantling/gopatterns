package main

import (
	"fmt"
)

// Observer observes changes in state
// Go is a smartass and if you create multiple instances of the same empty struct, there is only one object created at one memory address.
// So we add an unused field to guarantee unique addresses.
type Observer struct {
	unused int
}

// Observe is called when state changes
func (o *Observer) Observe(val string) {
	fmt.Printf("Observer at address %p observes value %q\n", o, val)
}

// Subject lets observers know of state changes
type Subject struct {
	observers []*Observer
}

// AddObserver adds and observer to the subject
func (s *Subject) AddObserver(observer *Observer) {
	s.observers = append(s.observers, observer)
}

// Update updates the observed value
func (s Subject) Update(val string) {
	fmt.Println("Updated value", val)

	for _, observer := range s.observers {
		observer.Observe(val)
	}
}

func main() {
	var subj Subject

	subj.AddObserver(&Observer{})
	subj.AddObserver(&Observer{})
	subj.AddObserver(&Observer{})

	subj.Update("first")
	subj.Update("second")
}
