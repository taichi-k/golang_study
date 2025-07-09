package main

import "fmt"

func adder() func(int) int {
	sum := 0 // sumはクロージャの外側で定義されている変数
	// クロージャはsumを参照できる
	// つまり、sumはクロージャの状態を保持する
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
