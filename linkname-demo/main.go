package main

import _ "unsafe"

//go:linkname nanotime runtime.nanotime
func nanotime() int64

func main() {
	println(nanotime())
}
