package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func B() string {
	_, file, line, _ := runtime.Caller(1)

	idx := strings.LastIndexByte(file, '/')

	return "=> " + file[idx+1:] + ":" + strconv.Itoa(line)
}

func A() string {
	return B()
}

func main() {
	test := "Here is $1 which is $2!"

	test = strings.ReplaceAll(test, "$1", "honey")
	test = strings.ReplaceAll(test, "$2", "tasty")
	fmt.Println(test)

	te := "aba abba abbba"
	re := regexp.MustCompile(`b+`)
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)

	fmt.Println(mm) // [b bb bbbb]
	fmt.Println(id) // [[1 2] [5 7] [10 13]]

	for _, d := range id {
		fmt.Println(te[d[0]:d[1]])
	}
	fmt.Println(A())
}
