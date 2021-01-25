// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"strconv"
	"time"
)

// Task is a task that takes some time to complete
type Task interface {
	Perform(string) int
}

// TaskCache is a proxy of Task that delegates to an underlying idempotent time consuming task.
// It remembers the results for known input values to return them immediately.
type TaskCache struct {
	task         Task
	knownResults map[string]int
}

// NewTaskCache constructs a TaskCache with a specific time consuming task
func NewTaskCache(timeConsuming Task) *TaskCache {
	return &TaskCache{
		task:         timeConsuming,
		knownResults: map[string]int{},
	}
}

// Perform only calculates the result for a given input once
func (c *TaskCache) Perform(input string) int {
	if result, exists := c.knownResults[input]; exists {
		return result
	}

	result := c.task.Perform(input)
	c.knownResults[input] = result
	return result
}

// TimeConsumingTask takes 1 second to execute
type TimeConsumingTask struct{}

// Perform takes 1 second
func (t TimeConsumingTask) Perform(input string) int {
	oneSec, _ := time.ParseDuration("1s")
	time.Sleep(oneSec)

	result, _ := strconv.Atoi(input)
	return result
}

func main() {
	proxy := NewTaskCache(TimeConsumingTask{})
	fmt.Printf("Value for 1 = ")
	fmt.Println(proxy.Perform("1"))
	fmt.Printf("Value for 1 = ")
	fmt.Println(proxy.Perform("1"))
	fmt.Printf("Value for 2 = ")
	fmt.Println(proxy.Perform("2"))
	fmt.Printf("Value for 2 = ")
	fmt.Println(proxy.Perform("2"))
}
