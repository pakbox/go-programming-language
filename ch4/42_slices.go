package main

import (
	"fmt"
)

// reverse reverses a slice of ints in place
func reverse(s []int) {
	for i, j := 1, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotate, left rotates a slice of ints by n positions
func rotate(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s[:])
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y ...int) []int {
	fmt.Printf("cap(x): %d, len(x): %d\n", cap(x), len(x))
	var z []int
	zlen := len(x) + len(y)
	fmt.Printf("zlen: %d\n", zlen)
	if zlen <= cap(x) {
		fmt.Println("--if--")
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		fmt.Println("--else--")
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			fmt.Printf("doubling\n")
			zcap = 2 * len(x)
		}
		fmt.Printf("zcap: %d\n\n", zcap)
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

// nonempty return a slice holding only the non-empty strings
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// remove1 - preserves order
func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// remove2 - dosen't preserves order
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Printf("Months: %v\n", months)
	fmt.Printf("len: %d, cap: %d\n", len(months), cap(months))
	Q2 := months[4:7]
	fmt.Printf("Q2 len: %d, cap: %d\n", len(Q2), cap(Q2))
	summer := months[6:9]
	fmt.Printf("summer len: %d, cap: %d\n", len(summer), cap(summer))
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	//fmt.Println(summer[:20]) // panic: out of range
	endlessSummer := summer[:5] // extend a slice, within capacity
	fmt.Println(endlessSummer)

	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	//reverse(intSlice)
	rotate(intSlice, 2)
	fmt.Println(intSlice)
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "b", "c"}
	s3 := []string{"a", "d", "c"}
	fmt.Printf("%v %v %t\n", s1, s2, equal(s1, s2))
	fmt.Printf("%v %v %t\n", s1, s3, equal(s1, s3))
	x1 := make([]int, 2, 4)
	fmt.Println(x1, len(x1), cap(x1))
	x1 = append(x1, 1, 2, 3)
	fmt.Println(x1, len(x1), cap(x1))

	var runes []rune
	pp := "helloworld"
	fmt.Printf("pp: %s, type: %T\n", pp, pp)
	for _, v := range "helloworld" {
		fmt.Printf("type: %T\n", v)
		runes = append(runes, v)
	}
	yy := []int{}
	/*
			yy1 := appendInt(yy, 3)
			yy2 := appendInt(yy1, 4)
			yy3 := appendInt(yy2, 5)
			fmt.Println(yy3)
		for i := 0; i < 5; i++ {
			yy = appendInt(yy, i, 4, 5)
		}
	*/
	fmt.Println(yy)

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2))
	fmt.Printf("%q\n", data2)

	type stack int
	st := []stack{}
	st = append(st, 1)
	st = append(st, 2)
	st = append(st, 3)
	top := st[len(st)-1]
	st = st[:len(st)-1]
	fmt.Println(st, top)

	rt := []int{5, 6, 7, 8, 9}
	fmt.Println(rt)
	rt = remove1(rt, 2)
	fmt.Println(rt)
	rt = remove2(rt, 2)
	fmt.Println(rt)
}
