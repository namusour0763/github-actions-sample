package main

import (
	"example/sqrt"
	"fmt"
)

func main() {
	fmt.Println("sqrt(2) =", sqrt.Sqrt(2))
	fmt.Println("sqrt(4) =", sqrt.Sqrt(4))
	fmt.Println("sqrt(-1) =", sqrt.Sqrt(-1))
}
