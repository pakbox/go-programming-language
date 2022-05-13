package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("p, type(p): %d, %T\n", *p, p)
}
