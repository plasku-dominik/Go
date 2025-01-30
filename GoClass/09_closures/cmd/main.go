// closures
package main

import (
	"fmt"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func do(d func()) {
	d()
}

func main() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}

	h, g := fib(), fib()

	fmt.Println(h(), h(), h(), h(), h())
	fmt.Println(g(), g(), g(), g(), g())

	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

		do(v)
	}

	s := make([]func(), 4)

	for j := 0; j < 4; j++ {
		j2 := j // closure capture, worked without it however
		s[j] = func() {
			fmt.Printf("%d @ %p\n", j2, &j2)
		}
	}

	for j := 0; j < 4; j++ {
		s[j]()
	}
}
