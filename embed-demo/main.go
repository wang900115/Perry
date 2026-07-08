package main

import (
	_ "embed"
	"fmt"
)

//go:embed assets/hello.txt
var content string

func main() {
	fmt.Println(content)
}
