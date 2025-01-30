// functions, parameters & defer
package main

import (
	"fmt"
)

func do(m1 *map[int]int) /*int*/ {
	(*m1)[3] = 0
	//fmt.Printf("b@ %p\n", b)
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("m1", *m1)
}

func main() {
	//a := []int{1, 2, 3}
	//fmt.Printf("a@ %p\n", a)
	m := map[int]int{4: 1, 7: 2, 8: 3}
	/*v :=* do(a)*/
	fmt.Println("m", m)
	do(&m)
	fmt.Println("m", m)

}
