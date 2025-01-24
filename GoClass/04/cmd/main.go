// Strings
package main

import (
	"fmt"
)

func main() {
	s := "élite"
	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))

	b := []byte(s)
	fmt.Printf("%8T %[1]v\n", b, len(b))
}
