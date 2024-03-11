package main

import (
	"Diffie-Hellman/internal/crypto"
	"fmt"
)

func main() {
	info := crypto.Ceaser("helloworld", 25)
	fmt.Println(info)
}
