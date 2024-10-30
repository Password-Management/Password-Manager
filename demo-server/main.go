package main

import (
	"demo-server/server"
	"fmt"
)

func main() {
	fmt.Println("When demo is requested the request comes here")
	server.Server()
}