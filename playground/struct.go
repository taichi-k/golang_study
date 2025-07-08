package main

import "fmt"

type Vertex struct {
	X int
	Y string
}

func main() {
	v := Vertex{1, "abc"}
	p := &v
	p.X = 1e9
	p.Y = "xyz"
	v.Y = "def"
	fmt.Println(v)
}
