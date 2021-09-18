package main

import (
	"fmt"

	"github.com/colemalphrus/golp/server"
)

func main() {
	fmt.Println("====================================")
	fmt.Println("====================================")
	fmt.Println("====RUNNING DEV SERVER ON :8080=====")
	fmt.Println("====================================")
	fmt.Println("====================================")
	fmt.Print("CMD C to quit  >")
	server.Serve(":8080")
}
