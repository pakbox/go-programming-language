package main

import (
	"fmt"
	"crypto/sha256"
)

func foo(arr [5]int){
	arr[0] = 100
	fmt.Println(arr)
}

func bar(parr *[5]int){
	parr[0] = 100
	fmt.Println(*parr)
}

func main(){
	var a [3]int // array of 3 integers
	fmt.Println(a[0]) // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])

	q := [...]int{1,2,3}
	fmt.Printf("%T\n", q)

	//q = [4]int{1, 2, 3, 4}

	type Currency int

	const (
		USD Currency = iota // U.S.A
		EUR // Europe
		GBP // England
		RMB // China
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println(RMB, symbol[RMB])

	s := [...]int{99: -1}
	fmt.Println(s)

	a1 := [2]int{1, 2}
	b1 := [...]int{1,2}
	c1 := [2]int{1,3}
	fmt.Println(a1 == b1, b1 == c1, c1 == a1)

	cr1 := sha256.Sum256([]byte("x"))
	cr2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n\n", cr1, cr2, cr1 == cr2, cr1)

	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	foo(arr)
	fmt.Println(arr)
	fmt.Println("----")
	fmt.Println(arr)
	bar(&arr)
	fmt.Println(arr)
}
