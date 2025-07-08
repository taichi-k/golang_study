package main

var arr []int = []int{10, 20, 30}

func main() {
	for idx, v := range arr {
		println(idx, v)
	}
}
