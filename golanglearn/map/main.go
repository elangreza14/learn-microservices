package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 12
	v := m["k1"]
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["k2"])
	fmt.Println(v)
	delete(m, "k1")
	a, prs := m["k1"]
	fmt.Println(a)
	fmt.Println(prs)

	n := map[string]int{"s": 3, "d": 3}
	fmt.Println(n)
	g := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", g)
}
