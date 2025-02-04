package main

import (
	"fmt"
)

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items {
		/*
			this version of go just creates a new slice instead of just reference so it doesn't reference
			the underlying array(items) -> causing all elements of a to point to the last 'item' in memory
			i := make([]byte, len(item))

			copy(i, item[:]) // make unique
		*/
		a = append(a, item[:])
	}

	fmt.Println("items:", items)
	fmt.Println("a:", a)
}

/*
	Any struct with a mutex must be passed by reference:

	type Employee struct {
		mu sync.Mutex
		Name string
		...
	}

	func do(emp *Employee) {
		emp.mu.lock()

		defer emp.mu.Unlock()
		...
	}

	Any small struct under 64 bytes probably should be copied:

	type Widget struct {
		ID 	  int
		Count int
	}

	func Expend(w Widget) Widget {
		w.Count--

		return w
	}
	Go routinely copies string & slice descriptors

	Stack allocation is more efficient
	Accessing a variable directly is more efficient than following a pointer
	Accessing a dense sequence of data is more efficient than sparse data (an array is faster than a linked list, etc.)

	Go would prefer to allocate on the stack, but sometimes can't:
		- a function returns a pointer to a local object
		- a local object is captured in a function closure
		- a pointer to a local object is sent via a channel
		- any object is assigned into an interface
		- any object whose size is variable at runtime (slices)

	The value returned by range is always a copy
	for i, thing := range things {
		... thing is a copy
	}

	Use the index to mutate the element:
	for i := range things {
		things[i].which = whatever
		...
	}

*/
