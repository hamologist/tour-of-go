package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := -1
	next := -1

	return func() int {
		if next == -1 {
			next = 1
			return 0
		} else if prev == -1 {
			prev = 0
			return 1
		}

		iterResult := prev + next
		prev = next
		next = iterResult
		return iterResult
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
