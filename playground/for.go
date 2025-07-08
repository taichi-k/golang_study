package main

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}

	sum := 0
	for sum < 10 {
		println("sum: ", sum)
		sum++
	}
}
