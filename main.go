package main

import (
	"fmt"
	"sync"
)

// Merge function
func merge(cs ...<-chan int) chan int {

	// Wait group
	var wg sync.WaitGroup

	result := make(chan int)

	// Add tasks for WaitGroup
	wg.Add(len(cs))

	// Loop through channels
	for _, c := range cs {

		// Gouroutine
		go func(c <-chan int) {
			defer wg.Done()
			for ch := range c {
				result <- ch
			}
		}(c)

	}
	// Close result channel
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

// Fill channel
func fillchan(num int) chan int {

	result := make(chan int)

	go func() {
		for i := 0; i < num; i++ {

			result <- i

		}
		close(result)
	}()

	return result
}

func main() {

	a := fillchan(4)

	b := fillchan(3)

	c := fillchan(2)

	d := merge(a, b, c)

	for i := range d {
		fmt.Println(i)
	}

}
