package main

import (
	"fmt"

	"github.com/colemalphrus/golp/config"
)

func main() {
	fmt.Println("====================================")
	fmt.Println("====================================")
	fmt.Println("====RUNNING DEV SERVER ON :8080=====")
	fmt.Println("====================================")
	fmt.Println("====================================")
	fmt.Print("CMD C to quit  >")
	config.Serve(":8080")
}
