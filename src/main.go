package main

import (
	"a"
	"b"
	"c/c1"
	"c/c2"
	"fmt"
)

func testA() {
	fmt.Println("test")
}
func main() {
	a.PrintA()
	b.PrintB1()
	b.PrintB2()
	c1.PrintC1()
	c2.PrintC2()
	testA()
}
