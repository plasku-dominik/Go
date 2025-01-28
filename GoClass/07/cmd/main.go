// Formatted & File I/O
package main

import (
	"fmt"
)

func main() {
	// string format
	a, b := 12, 345
	c, d := 1.2, 3.45

	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%X %X\n", a, b)   // x,X <- base-16, upper case
	fmt.Printf("%#x %#x\n", a, b) // # <- 0x before the number
	fmt.Printf("%f %.3f\n", c, d) // .? before the f <- 2 decimal

	fmt.Println()

	fmt.Printf("|%9d|%9d|\n", a, b)   // column with the width of the character
	fmt.Printf("|%09d|%09d|\n", a, b) // filling up the space with 0
	fmt.Printf("|%-9d|%-9d|\n", a, b) // start of the column
	fmt.Printf("|%9f|%9.2f|\n", c, d)

	// slices
	fmt.Println()
	s := []int{1, 2, 3}
	t := [3]rune{'a', 'b', 'c'}

	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

	fmt.Println()
	fmt.Printf("%T\n", t)
	fmt.Printf("%q\n", t)
	fmt.Printf("%v\n", t)
	fmt.Printf("%#v\n", t)

	// map
	m := map[string]int{"and": 1, "or": 2}
	s2 := "a string"
	b2 := []byte(s2)

	fmt.Println()
	fmt.Printf("%T\n", m)
	fmt.Printf("%q\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)

	fmt.Println()
	fmt.Printf("%T\n", s2)
	fmt.Printf("%q\n", s2)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%#v\n", s2)
	fmt.Printf("%v\n", string(b2))

}
