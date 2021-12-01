package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["juan"] = 29
	m1["celina"] = 66
	fmt.Println(m1["celina"])
}
