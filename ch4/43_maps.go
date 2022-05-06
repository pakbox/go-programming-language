package main

import (
//	"bufio"
	"fmt"
//	"os"
	"sort"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok := y[xk]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func badequal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv := y[xk]; yv != xv {
			return false
		}
	}
	return true
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func main() {
	ages1 := make(map[string]int) // mapping from string to int
	ages1["alice"] = 31
	ages1["charlie"] = 34
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Printf("ages1: %v\n", ages1)
	fmt.Printf("ages2: %v\n", ages2)
	ages1["alice"] = 32
	fmt.Printf("ages1: %v\n", ages1)
	fmt.Printf("alice's age: %v\n", ages1["alice"])
	fmt.Printf("bob's age: %v\n", ages1["bob"])
	delete(ages2, "charlie")
	delete(ages2, "bob")
	fmt.Printf("ages2: %v\n", ages2)
	for name, age := range ages1 {
		fmt.Printf("%s\t%d\n", name, age)
	}
	//var names []string
	names := make([]string, 0, len(ages1))
	for name := range ages1 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages1[name])
	}
	age, ok := ages1["bob"]
	if !ok {
		fmt.Println("no key named bob in map ages1")
	} else {
		fmt.Println("age: ", age)
	}

	// alternate way of writing above logic
	if age, ok := ages1["bob"]; !ok {
		fmt.Println("no key named bob in map ages1")
	} else {
		fmt.Println("age: ", age)
	}

	// comparing maps
	m1 := map[string]int{"a": 0}
	m2 := map[string]int{"b": 4}
	fmt.Println(badequal(m1, m2), equal(m1, m2))

	// go does not have set type
	// emulating sets
	/*
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
	*/
	sl := []string{"a", "b", "c"}
	Add(sl)
	Add(sl)
	Add(sl)
	Add(sl)
	fmt.Println(Count(sl))
}
