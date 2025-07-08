package main

import "fmt"

var i, j = 1, 2

func main() {
	var a, b, c = true, 0, "no"

	fmt.Println("i:", i, "j:", j, "a:", a, "b:", b, "c:", c)
	fmt.Printf("i の型は %T\n", i)
	fmt.Printf("j の型は %T\n", j)
	fmt.Printf("a の型は %T\n", a)
	fmt.Printf("b の型は %T\n", b)
	fmt.Printf("c の型は %T\n", c)
}
