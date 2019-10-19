package main

import (
	"fmt"

	"github.com/PaesslerAG/gval"
)

func main() {
	fmt.Println("Hello World!")

	s := "1 + 4 * 2"

	v, err := gval.Evaluate(s, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
