package main

func main() {
	defer println("later")

	println("now")
}
