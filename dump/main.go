package main

import (
	"fmt"
	"time"
)

func main() {

	secret := "prod-db-password-123"

	fmt.Println("service started")

	fmt.Println(secret)

	time.Sleep(time.Hour)
}
