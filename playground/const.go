package main

import "fmt"

func main() {
	const Hello = "Hello, World!"
	const Pi = 3.14
	const MaxInt = 1<<63 - 1 // 最大のint値
	const IsTrue = true
	const (
		Red   = iota // 0
		Green        // 1
		Blue         // 2
	)
	const (
		Monday    = iota + 1 // 1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)
	const (
		_ = iota             // 0をスキップ
		A = 1 << (iota * 10) // 1024
		B = 1 << (iota * 10) // 1024 *
		C = 1 << (iota * 10) // 1024 * 1024
	)
	fmt.Println(Hello)
	fmt.Println(Pi)
	fmt.Println(MaxInt)
	fmt.Println(IsTrue)
	fmt.Println(Red, Green, Blue)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Println(A, B, C)
}
