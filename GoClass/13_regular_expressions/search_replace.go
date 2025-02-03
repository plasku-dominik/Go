package main

import (
	"fmt"
	"regexp"
)

var ph = regexp.MustCompile(`\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`)

func main() {
	orig := "(214) 514-9548"
	match := ph.FindStringSubmatch(orig)

	fmt.Printf("%q\n", match)

	if len(match) > 3 {
		fmt.Printf("+1 %s-%s-%s\n", match[1], match[2], match[3])
	}
}
