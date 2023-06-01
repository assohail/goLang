package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := 0
	for a < len(os.Args) {
		fmt.Printf("a[%d] = %s\n", a, os.Args[a])
		a++
	}

	a5 := [5]float64{3, 5., 0.2}
	fmt.Printf("a5= %v\n", a5)

	view := a5[2:]
	fmt.Printf("view %v\n", view)

	view[0] = 5.
	fmt.Printf("a5 edit, %v\n", a5)
	fmt.Printf("view edit %v\n", view)

	a10 := [10]float64{64, 63, 62}
	view = append(view, a10[:]...)

	fmt.Printf("view a10 %v\n", view)

	// strings
	s := "cat is white"
	cat := s + ". with voice mhewon!"

	for k, v := range cat {
		fmt.Printf("[%d] %c\n", k, v)
	}

	j := strings.Join([]string{"CATS", "DOGS", "COWS"}, ",")
	println(j)
	k := strings.Split(j, ",")
	fmt.Printf("resplit: = %v\n", k)

	// maps
	emap := make(map[string]int)
	fmt.Printf("%v\n", emap)

	m := map[string]int{"k": 5}
	m["p"] = 8
	m["v"] = 11
	m["no"] = 13

	delete(m, "v")

	fmt.Printf("m=%v\n", m)

	// closure
	sub := func(n int) int {
		return n - 7
	}

	fmt.Printf("sub(10) = %v\n", sub(10))

	a3 := AddNer(3)
	fmt.Printf("a3(6) %d\n", a3(6))
	fmt.Printf("a3(9) %d\n", a3(9))
}

func AddNer(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
