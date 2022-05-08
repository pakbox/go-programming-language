package main

import (
	"fmt"
)

type orange int
type apple int

type emp struct {
	id	int
	name	string
}

type dup struct {
	name	string
	id	int
}

var t dup
var T dup

func main() {
	var o1 orange
	var a1 apple
	o1 = 2
	a1 = 2
	_ = a1
	_ = o1
	// do not compare apple to oranges
	// invalid operation: a1 == o1 (mismatched types apple and orange)
	/*
	if a1 == o1 {
		fmt.Println("ok")
	}
	*/
	var e1 emp
	var d1 dup
	_ = e1
	_ = d1
	//fmt.Println(e1, d1, e1 == d1)
	var x struct{x, y int}
	var y struct{x, y int}
	x.x, x.y = 1, 2
	y.x, y.y = 1, 2
	fmt.Println(x == y)
}
