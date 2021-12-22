package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(sayHi())

	tesla := Car{
		name: "Tesla",
	}

	log.Println(tesla.Start())
	log.Println(tesla.Stop())
	log.Println(tesla.n)
}

func sayHi() string {
	return "Hi!"
}
