package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

type album1 struct {
	title string
}

type album2 struct {
	title string
}

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words"` //omitempty"`
}

func main() {
	fmt.Println("Employee struct:")
	c := map[string]*Employee{}

	c["Lamine"] = &Employee{"Lamine", 2, nil, time.Now()}
	c["Lamine"].Number++
	c["Matt"] = &Employee{
		Name:   "Matt",
		Number: 1,
		Boss:   c["Lamine"],
		Hired:  time.Now(),
	}

	fmt.Printf("%T %+[1]v\n", c["Matt"])
	fmt.Printf("%T %+[1]v\n", c["Lamine"])
	fmt.Println("\nAlbum example:")
	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The White Album",
		"The Beatles",
		1968,
		10000000000,
	}

	var pAlbum *struct { // album = struct works instead of *
		title  string
		artist string
		year   int
		copies int
	}
	fmt.Println(album, pAlbum)

	var a1 = album1{
		"The White Album",
	}

	var a2 = album2{
		"The Black Album",
	}

	a1 = album1(a2)
	fmt.Println(a1, a2)

	r := Response{Page: 1} //, Words: []string{"up", "in", "out"}}
	j, _ := json.Marshal(r)
	fmt.Println("\nJSON Example:\n", string(j))
	fmt.Printf("%#v\n", r)

	var r2 Response

	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)
}

/*
	Two struct types are the compatible if:
		- the fields have the same types and names
		- in the same order
		- and the same tags (*)

	A struct may be copied and passed as a parameter in its entirety
	A struct is comparable if all its fields are comparable
	The zero value for a struct is "zero" for each field in turn
*/
