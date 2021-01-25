// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
)

// StateTransition describes a transition from one state to another
type StateTransition interface {
	Transition() StateTransition
}

// abstractStatetTransition unconditionally transitions to one specific next state
type abstractStateTransition struct {
	nextState StateTransition
}

// Transition transitions to next state
func (s abstractStateTransition) Transition() StateTransition {
	return s.nextState
}

// EnterCodeState waits for user to enter the code to start washing their car
type EnterCodeState struct {
	abstractStateTransition
}

// NewEnterCodeState constructs an EnterCodeState
func NewEnterCodeState() *EnterCodeState {
	return &EnterCodeState{
		abstractStateTransition{
			nextState: NewOpenEntranceDoorState(),
		},
	}
}

// OpenEntranceDoorState opens the door so the user can drive their car into the wash
type OpenEntranceDoorState struct {
	abstractStateTransition
}

// NewOpenEntranceDoorState constructs an OpenEntranceDoorState
func NewOpenEntranceDoorState() *OpenEntranceDoorState {
	return &OpenEntranceDoorState{
		abstractStateTransition{
			nextState: NewCloseEntranceDoorState(),
		},
	}
}

// CloseEntranceDoorState closes the door after the user enters the wash
type CloseEntranceDoorState struct {
	abstractStateTransition
}

// NewCloseEntranceDoorState constructs a CloseEntranceDoorState
func NewCloseEntranceDoorState() *CloseEntranceDoorState {
	return &CloseEntranceDoorState{
		abstractStateTransition{
			nextState: NewWashState(),
		},
	}
}

// WashState washes the car
type WashState struct {
	abstractStateTransition
}

// NewWashState constructs a WashState
func NewWashState() *WashState {
	return &WashState{
		abstractStateTransition{
			nextState: NewRinseState(),
		},
	}
}

// RinseState rinses the car
type RinseState struct {
	abstractStateTransition
}

// NewRinseState constructs a RinseState
func NewRinseState() *RinseState {
	return &RinseState{
		abstractStateTransition{
			nextState: NewOpenExitDoorState(),
		},
	}
}

// OpenExitDoorState opens the door so the user can drive their car into the wash
type OpenExitDoorState struct {
	abstractStateTransition
}

// NewOpenExitDoorState constructs an OpenExitDoorState
func NewOpenExitDoorState() *OpenExitDoorState {
	return &OpenExitDoorState{
		abstractStateTransition{
			nextState: NewCloseExitDoorState(),
		},
	}
}

// CloseExitDoorState closes the door after the user enters the wash
type CloseExitDoorState struct {
	abstractStateTransition
}

// NewCloseExitDoorState constructs an CloseExitDoorState
func NewCloseExitDoorState() *CloseExitDoorState {
	return &CloseExitDoorState{
		abstractStateTransition{
			nextState: nil,
		},
	}
}

// CarWashContext maintains the context of a car wash
type CarWashContext struct {
	currentState StateTransition
}

// NewCarWashContext constructs a CarWashContext
func NewCarWashContext() *CarWashContext {
	return &CarWashContext{
		currentState: NewEnterCodeState(),
	}
}

// Transition transitions to next state
func (w *CarWashContext) Transition() StateTransition {
	curState := w.currentState
	nextState := w.currentState.Transition()
	w.currentState = nextState

	fmt.Printf("Transition from %T to %T\n", curState, nextState)
	return nextState
}

func main() {
	w := NewCarWashContext()
	for {
		nextState := w.Transition()
		if nextState == nil {
			break
		}
	}
}
