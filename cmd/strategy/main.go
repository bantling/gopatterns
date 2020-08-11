package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// bubbleSort sorts at least one item each pass, takes n^2 time
func bubbleSort(items []int) {
	// Swap backward, same direction as iteration, to guarantee at last one item bubbles to correct location each time.
	// As such, we iterate one less item each pass for a bit of extra speed.
	for i := 0; i < len(items) - 1; i++ {
		for k := len(items) - 1; k > i; k-- {
			if items[k - 1] > items[k] {
				tmp := items[k - 1]
				items[k - 1] = items[k]
				items[k] = tmp
			}
		}
	}
}

// goSort sorts items using Go sort api
func goSort(items []int) {
	sort.Ints(items)
}

// doSort chooses an algorithm based on the number of items to sort
func doSort(items []int) {
	if len(items) <= 10 {
		fmt.Printf("bubbleSort: %v = ", items)
		bubbleSort(items)
	} else {
		fmt.Printf("golSort: %v = ", items)
		goSort(items)
	}
	
	fmt.Printf("%v\n", items)
}

func main() {
	list := []int{2,1,3,5,4,6}
	bubbleSort(list)
	fmt.Printf("%v\n", list)
	
	list = []int{2,1,3,5,4,6}
	goSort(list)
	fmt.Printf("%v\n", list)
	
	list = make([]int, 10)
	for i := 0; i < len(list); i++ {
		list[i] = rand.Int() % 100
	}
	doSort(list)
	
	list = make([]int, 15)
	for i := 0; i < len(list); i++ {
		list[i] = rand.Int() % 100
	}
	doSort(list)
}
