// slices: describing, length, capacity, appends
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s []int

	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	// %d -> base 10 | %T -> type of value | %5t -> true or false, 5 means min length of output | %#[?]v -> a Go-syntax representation of the value, ? -> ?.argument
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil) // 1st one is nil
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil) // 2nd one is empty
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil) // 5 elements | 5 capacity | []int slice | not nil | contains five 0
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil) // 0 elements in slice | 5 max capacity | [] int slice | v != nil | shows empty slice

	var x []int

	jl, _ := json.Marshal(x)
	fmt.Println(string(jl)) // nil slice -> null

	y := []int{}

	j2, _ := json.Marshal(y)
	fmt.Println(string(j2)) // empty slice -> empty

	a := [...]int{1, 2, 3}
	//b := a[:1]
	b := []int{1, 2, 3}
	c := b[0:2:2]
	//d := a[0:1:1] // [i:j:k] len j-i | cap k-i -> length of 1, capacity 1

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	fmt.Println("b len:", len(b))
	fmt.Println("b capacity:", cap(b))

	fmt.Println("c =", c)
	fmt.Println("c len:", len(c))
	fmt.Println("c capacity:", cap(c))
	/*
		fmt.Println("d =", d)
		fmt.Println("d len:", len(d))
		fmt.Println("d capacity:", cap(d))
	*/

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c[0] = 9
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
}
