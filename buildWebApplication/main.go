package main

import (
	"service"
)

func main() {
	server := service.NewServer()
	server.Run(":8080")
}
