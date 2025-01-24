// Basic Types
package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		a := 2
		b := 3.1

		fmt.Printf("a: %8T %[1]v\n", a, a) // %8T -> allign + type | %v -> value | [1] -> reuse w/e parameter 1 was
		fmt.Printf("b: %8T %[1]v\n", b, b)

		a = int(b)
		fmt.Printf("a: %8T %v\n", a, a)
		b = float64(a)
		fmt.Printf("b: %8T %v\n", b, b)
	*/
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val) // this can be put between the if and err to have the same effect, put ; before err in this case
		if err != nil {
			break
		}

		sum += val // sum = sum + val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(n))
}
