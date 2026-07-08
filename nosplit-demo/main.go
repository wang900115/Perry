package main

//go:nosplit
func hello() {
	println("Hello")
}

//go:nosplit
func ba() {
	var buf [10000]int
	println(buf[0])
}

func main() {
	hello()
	ba()
}
