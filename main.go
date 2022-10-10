package main

import (
	_ "dtskominfo-gin/database"
	"dtskominfo-gin/routers"
	"fmt"
)

func main() {
	fmt.Println("Start Server ")
	routers.GetRouter().Run(":8000")
}