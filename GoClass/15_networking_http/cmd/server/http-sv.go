package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const url = "http://jsonplaceholder.typicode.com"

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body) // ioutil is otudated since 1.16 so using io

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
