package main

import (
	"fmt"
	_ "rajastra/rest-api-go/database"
	"time"
)

func main() {
	fmt.Println("Test auto migrate")
	time.Sleep(time.Duration(5) * time.Second)
}