package main

import (
	"fmt"
	"regexp"
)

const phre = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`

var pfmt = regexp.MustCompile(phre)

func main() {
	r, b := pfmt.LiteralPrefix()
	fmt.Printf("%q %t\n", r, b)

	orig := "call me at (214) 514-3232 today"

	match := pfmt.FindStringSubmatch(orig)
	fmt.Printf("%q\n", match)

	intl := pfmt.ReplaceAllString(orig, "+1 ${1}-${2}-${3}")

	fmt.Println(intl)
}
