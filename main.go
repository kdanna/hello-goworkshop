package main

import (
	"fmt"

	u "github.com/mystringutils"
)

func main() {
	s := "Hello, World"
	fmt.Println("Hello, World")
	fmt.Println(u.Lower(s))
	fmt.Println(u.Upper(s))
}
