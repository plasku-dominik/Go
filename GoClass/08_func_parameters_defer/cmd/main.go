// functions, parameters & defer
package main

import (
	"fmt"
)

func doIt() (a int) {
	defer func() {
		a = 2
	}()

	a = 1
	return
	// first stack runs after we exit with the last return
}

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

	a := 10
	defer fmt.Println("original a value:", a)
	a = 11
	fmt.Println("modified a:", a)

	fmt.Println("do it func:", doIt())

}

/*
	Recursion example
	func walk(node *tree.T) int {
		if node == nil {
			return 0
	}

	return node.value + walk(node.left) + walk(node.right)
}
*/

/*
	Defer example
	func main() {
		f, err := os.Open("my_file.txt")

		if err != nil {
			...
		}
		defer f.Close()

		and do something with the file
	}

	func main() {
		f := os.Stdin

		if len(os.Args) > 1 {
			if f, err := os.Open(os.Args[1]); err != nil {
				...
			}
			defer f.Close()
		}
		file still usable, only closes when the function exists
	}

	The deferred calls to Close must wait until function exit (might run out of file descriptors before that)
	func main() {
		for i := 1; i < len(os.Args); i++ {
			f, err := os.Open(os.Args[i])
			...
			defer f.Close()
			...
		}
	}
*/
