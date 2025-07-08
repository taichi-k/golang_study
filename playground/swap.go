package main

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	var a, b string = swap("hello", "world")
	println(a, b)
}
